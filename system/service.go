package system

import (
	"fmt"
	"go-systemd-docker/utils"
	"os"
	"path"

	"github.com/kardianos/service"
	"gopkg.in/yaml.v3"
)

var indexpath string = path.Join(utils.CONFIG_DIR_PATH, utils.INDEX_FILE_NAME_WITH_EXT)

// CreateService create a new service config and save it as a file & also update index.
func CreateService(sys *System) (*service.Config, error) {
	byteYaml, err := yaml.Marshal(sys)
	if err != nil {
		return nil, err
	}

	if err := os.MkdirAll(utils.MANIFEST_DIR_PATH, utils.MANIFEST_DIR_PERM); err != nil {
		return nil, err
	}

	filepath := path.Join(utils.MANIFEST_DIR_PATH, sys.Name+utils.YAML_EXT)
	file, err := os.Create(filepath)
	if err != nil {
		return nil, err
	}

	if _, err = file.Write(byteYaml); err != nil {
		return nil, err
	}

	var byteIndex []byte

	// create index.yaml if not present with default configuration.
	if _, err := os.Stat(indexpath); os.IsNotExist(err) {
		if byteIndex, err = yaml.Marshal(&index{
			Name:     utils.INDEX_FILE_NAME_WITHOUT_EXT,
			Services: []indexService{},
		}); err != nil {
			return nil, err
		}

		// if err = os.WriteFile(indexpath, byteIndex, utils.MANIFEST_FILE_PERM); err != nil {
		// 	return nil, err
		// }
	}

	if byteIndex == nil {
		if byteIndex, err = os.ReadFile(indexpath); err != nil {
			return nil, err
		}
	}

	index := index{}
	if err = yaml.Unmarshal(byteIndex, &index); err != nil {
		return nil, err
	}

	// append in index.services the newly create services.
	index.Services = append(index.Services, indexService{
		Name: sys.Name,
		Path: filepath,
	})

	if byteIndex, err = yaml.Marshal(&index); err != nil {
		return nil, err
	}

	if err = os.WriteFile(indexpath, byteIndex, utils.MANIFEST_FILE_PERM); err != nil {
		return nil, err
	}

	return &service.Config{
		Name:        sys.Name,
		DisplayName: sys.DisplayName,
		Description: sys.Description,
		Executable:  sys.Executable,
		Arguments:   sys.Arguments,
	}, nil
}

// DeleteService deletes the service config created as a file & also update index.
func DeleteService(instanceName string) error {
	var err error
	var errs []error

	// removes manifest file.
	if !IsServiceExist(instanceName) {
		errs = append(errs, fmt.Errorf("manifest file not found"))
	}
	err = os.Remove(path.Join(utils.MANIFEST_DIR_PATH, instanceName+utils.YAML_EXT))

	if _, err := os.Stat(indexpath); os.IsNotExist(err) {
		errs = append(errs, fmt.Errorf("index file not found"))
		return fmt.Errorf("%v", errs)
	}

	byteIndex, err := os.ReadFile(indexpath)
	if err != nil {
		errs = append(errs, err)
		return fmt.Errorf("%v", errs)
	}

	index := index{}
	if err = yaml.Unmarshal(byteIndex, &index); err != nil {
		errs = append(errs, err)
		return fmt.Errorf("%v", errs)
	}

	for i, element := range index.Services {
		if element.Name == instanceName {
			index.Services = append(index.Services[:i], index.Services[i+1:]...)
		}
	}

	if byteIndex, err = yaml.Marshal(&index); err != nil {
		errs = append(errs, err)
		return fmt.Errorf("%v", errs)
	}

	if err = os.WriteFile(indexpath, byteIndex, utils.MANIFEST_FILE_PERM); err != nil {
		errs = append(errs, err)
		return fmt.Errorf("%v", errs)
	}

	return nil
}

// GetService return saved service Config as file.
func GetService(instanceName string) (*service.Config, error) {

	if !IsServiceExist(instanceName) {
		return nil, fmt.Errorf("service doesn't exist")
	}

	byteCode, err := os.ReadFile(path.Join(utils.MANIFEST_DIR_PATH, instanceName+utils.YAML_EXT))
	if err != nil {
		return nil, err
	}

	sys := System{}
	if err = yaml.Unmarshal(byteCode, &sys); err != nil {
		return nil, err
	}

	return &service.Config{
		Name:        sys.Name,
		DisplayName: sys.DisplayName,
		Description: sys.Description,
		Executable:  sys.Executable,
		Arguments:   sys.Arguments,
	}, nil
}

// ListServices return saved service.Config{} as file.
func ListServices() ([]indexService, error) {

	if _, err := os.Stat(indexpath); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("No service exist")
		}
		return nil, err
	}

	byteCode, err := os.ReadFile(indexpath)
	if err != nil {
		return nil, err
	}

	index := index{}
	if err = yaml.Unmarshal(byteCode, &index); err != nil {
		return nil, err
	}

	return index.Services, nil
}

// IsServiceExist checks if service exist.
func IsServiceExist(instanceName string) bool {
	if _, err := os.Stat(path.Join(utils.MANIFEST_DIR_PATH, instanceName+utils.YAML_EXT)); err != nil {
		return false
	}

	return true
}

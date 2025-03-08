package system

import (
	"fmt"
	"go-systemd-docker/utils"
	"os"
	"path"

	"github.com/kardianos/service"
	"gopkg.in/yaml.v3"
)

// Create a new service.Config{} and save as file
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
	indexpath := path.Join(utils.CONFIG_DIR_PATH, utils.INDEX_FILE_NAME_WITH_EXT)

	// create index.yaml if not present with default configuration.
	if _, err := os.Stat(indexpath); os.IsNotExist(err) {
		byteIndex, err = yaml.Marshal(index{
			name:     utils.INDEX_FILE_NAME_WITHOUT_EXT,
			services: []indexService{},
		})
		if err != nil {
			return nil, err
		}

		fmt.Println("index contents:", string(byteIndex))
		if err = os.WriteFile(indexpath, byteIndex, utils.MANIFEST_FILE_PERM); err != nil {
			return nil, err
		}
	}

	// if byteIndex == nil {
	// 	if byteIndex, err = os.ReadFile(indexpath); err != nil {
	// 		return nil, err
	// 	}
	// }

	// index := index{}
	// if err = yaml.Unmarshal(byteIndex, &index); err != nil {
	// 	return nil, err
	// }

	// // append in index.services the newly create services.
	// index.services = append(index.services, indexService{
	// 	name: sys.Name,
	// 	path: filepath,
	// })

	// if err = os.WriteFile(indexpath, byteIndex, utils.MANIFEST_FILE_PERM); err != nil {
	// 	return nil, err
	// }

	return &service.Config{
		Name:        sys.Name,
		DisplayName: sys.DisplayName,
		Description: sys.Description,
		Executable:  sys.Executable,
		Arguments:   sys.Arguments,
	}, nil
}

// DeleteService deletes the saved service.Config{} as file.
func DeleteService(instanceName string) error {
	var err error
	if IsServiceExist(instanceName) {
		err = os.Remove(path.Join(utils.MANIFEST_DIR_PATH, instanceName+utils.YAML_EXT))
	} else {
		err = fmt.Errorf("manifest file not found")
	}
	return err
}

// GetService return saved service.Config{} as file.
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
func ListServices() ([]string, error) {

	// if !IsServiceExist(instanceName) {
	// 	return nil, fmt.Errorf("service doesn't exist")
	// }

	// byteCode, err := os.ReadFile(path.Join(utils.MANIFEST_DIR_PATH, instanceName+utils.YAML_EXT))
	// if err != nil {
	// 	return nil, err
	// }

	// sys := System{}
	// if err = yaml.Unmarshal(byteCode, &sys); err != nil {
	// 	return nil, err
	// }

	return []string{}, nil
}

// IsServiceExist checks if service exist.
func IsServiceExist(instanceName string) bool {
	if _, err := os.Stat(path.Join(utils.MANIFEST_DIR_PATH, instanceName+utils.YAML_EXT)); err != nil {
		return false
	}

	return true
}

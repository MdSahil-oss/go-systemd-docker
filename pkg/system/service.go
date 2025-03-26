package system

import (
	"fmt"
	"go-systemd-docker/pkg/utils"
	"os"
	"path"

	"github.com/kardianos/service"
	"gopkg.in/yaml.v3"
)

// CreateService create a new service config and save it as a file & also update index.
func CreateService(sys *System, imageName string) (*service.Config, error) {
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
	if _, err := os.Stat(utils.INDEX_FILE_PATH); os.IsNotExist(err) {
		if byteIndex, err = yaml.Marshal(NewIndex(
			withIndexName(utils.INDEX_FILE_NAME_WITHOUT_EXT),
			withIndexServices([]IndexService{}),
		)); err != nil {
			return nil, err
		}
	}

	if byteIndex == nil {
		if byteIndex, err = os.ReadFile(utils.INDEX_FILE_PATH); err != nil {
			return nil, err
		}
	}

	index := Index{}
	if err = yaml.Unmarshal(byteIndex, &index); err != nil {
		return nil, err
	}

	// append in index.services the newly create services.
	index.Services = append(index.Services, *NewIndexService(
		withIndexServiceName(sys.Name),
		withIndexServicePath(filepath),
		withIndexServiceImage(imageName),
	))

	if byteIndex, err = yaml.Marshal(&index); err != nil {
		return nil, err
	}

	if err = os.WriteFile(utils.INDEX_FILE_PATH, byteIndex, utils.MANIFEST_FILE_PERM); err != nil {
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

	if _, err := os.Stat(utils.INDEX_FILE_PATH); os.IsNotExist(err) {
		errs = append(errs, fmt.Errorf("index file not found"))
		return fmt.Errorf("%v", errs)
	}

	byteIndex, err := os.ReadFile(utils.INDEX_FILE_PATH)
	if err != nil {
		errs = append(errs, err)
		return fmt.Errorf("%v", errs)
	}

	index := Index{}
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

	if err = os.WriteFile(utils.INDEX_FILE_PATH, byteIndex, utils.MANIFEST_FILE_PERM); err != nil {
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

// ListService return the service with given name created as systemd instance.
func ListService(instanceName string) (IndexService, error) {

	svcs, err := ListServices()
	if err != nil {
		return IndexService{}, err
	}

	for _, svc := range svcs {
		if svc.Name == instanceName {
			return svc, nil
		}
	}

	return IndexService{}, fmt.Errorf("no systemd service exist with %s name", instanceName)
}

// GetSystemDProcess returns the installed systemd config `instanceName` on the system.
func GetSystemDProcess(instanceName string) (service.Service, error) {
	var s service.Service
	svcConfig, err := GetService(instanceName)
	if err != nil {
		return s, err
	}

	prg := &CreateProgram{}
	s, err = service.New(prg, svcConfig)

	return s, err
}

// GetSystemDProcesses returns a list of installed systemd configs on the system.
func GetSystemDProcesses() ([]service.Service, error) {
	var svcs []service.Service

	ss, err := ListServices()
	if err != nil {
		return nil, err
	}

	var errs []error = nil
	for _, s := range ss {
		svc, err := GetSystemDProcess(s.Name)
		if err != nil {
			errs = append(errs, err)
		}

		svcs = append(svcs, svc)
	}

	if len(errs) > 0 {
		return svcs, fmt.Errorf("%v", errs)
	}
	return svcs, nil
}

// ListServices return a list of services created as systemd instance.
func ListServices() ([]IndexService, error) {

	if _, err := os.Stat(utils.INDEX_FILE_PATH); err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("No service exist")
		}
		return nil, err
	}

	byteCode, err := os.ReadFile(utils.INDEX_FILE_PATH)
	if err != nil {
		return nil, err
	}

	index := Index{}
	if err = yaml.Unmarshal(byteCode, &index); err != nil {
		return nil, err
	}

	return index.Services, nil
}

// ListRunningService return a list of systemd running services
func ListRunningService(instanceName string) (IndexService, error) {
	is, err := ListService(instanceName)
	if err != nil {
		return IndexService{}, err
	}

	svc, err := GetSystemDProcess(instanceName)
	if err != nil {
		return IndexService{}, err
	}

	status, err := svc.Status()
	if err != nil {
		return IndexService{}, err
	}

	if status == service.StatusRunning {
		is.Status = utils.PROCESS_STATUS_RUNNING
		return is, nil
	}

	return IndexService{}, err
}

// ListRunningService return a list of systemd running services
func ListRunningServices() ([]IndexService, error) {
	var indSvcs []IndexService

	ss, err := ListServices()
	if err != nil {
		return nil, err
	}

	indSvcs = ss
	var runningSS []IndexService
	var errs []error = nil
	for _, s := range indSvcs {
		svc, err := GetSystemDProcess(s.Name)
		if err != nil {
			errs = append(errs, err)
		}

		status, err := svc.Status()
		if err != nil {
			errs = append(errs, err)
		}

		if status == service.StatusRunning {
			s.Status = utils.PROCESS_STATUS_RUNNING
			runningSS = append(runningSS, s)
		}
	}

	if len(errs) > 0 {
		return runningSS, fmt.Errorf(fmt.Sprintf("%v", errs))
	}

	return runningSS, nil
}

// IsServiceExist checks if service exist.
func IsServiceExist(instanceName string) bool {
	if _, err := os.Stat(path.Join(utils.MANIFEST_DIR_PATH, instanceName+utils.YAML_EXT)); err != nil {
		return false
	}

	return true
}

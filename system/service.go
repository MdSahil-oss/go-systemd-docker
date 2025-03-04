package system

import (
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

	return &service.Config{
		Name:        sys.Name,
		DisplayName: sys.DisplayName,
		Description: sys.Description,
		Executable:  sys.Executable,
		Arguments:   sys.Arguments,
	}, nil
}

func IsServiceExist(instanceName string) bool {
	contents, err := os.ReadDir(utils.MANIFEST_DIR_PATH)
	if err != nil {
		return false
	}

	for _, content := range contents {
		if content.Name() == instanceName+utils.YAML_EXT {
			return true
		}
	}

	return false
}

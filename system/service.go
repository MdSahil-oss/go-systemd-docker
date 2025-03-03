package system

import (
	"go-systemd-docker/utils"
	"os"
	"path"

	"github.com/kardianos/service"
	"gopkg.in/yaml.v3"
)

// Create -> Check if the instance already exist
// If not then Calls NewSystem() and create one follows by CreateService

// Create a new service.Config{} and save as file
func CreateService(sys *System) (*service.Config, error) {
	byteYaml, err := yaml.Marshal(sys)
	if err != nil {
		return nil, err
	}

	if err = os.WriteFile(
		path.Join(utils.CONFIG_DIR_PATH, sys.Name+".yml"),
		byteYaml,
		0644,
	); err != nil {
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

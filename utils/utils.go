package utils

import (
	"fmt"
	"os"
	"path"
)

const (
	CONFIG_DIR                     = ".sysd"
	MANIFEST_FILE_PERM             = 0644
	MANIFEST_DIR_PERM              = 0775
	DEFAULT_DOCKER_EXECUTABLE_PATH = "/usr/bin/docker"
	YAML_EXT                       = ".yaml"
)

var (
	CONFIG_DIR_PATH   = path.Join(GetHomeDir(), CONFIG_DIR)
	MANIFEST_DIR_PATH = path.Join(CONFIG_DIR_PATH, "manifests")
)

// GetDockerExecutablePath returns executable docker path
func GetDockerExecutablePath() string {
	return DEFAULT_DOCKER_EXECUTABLE_PATH
}

func GetHomeDir() string {
	str, err := os.UserHomeDir()
	if err != nil {
		return "~/"
	}

	return str
}

// Terminate prints given string and exit with 1
func Terminate(str string) {
	fmt.Println(str)
	os.Exit(1)
}

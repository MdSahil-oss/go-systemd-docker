package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

const (
	CONFIG_DIR                     = ".sysd"
	MANIFEST_FILE_PERM             = 0644
	MANIFEST_DIR_PERM              = 0775
	DEFAULT_DOCKER_EXECUTABLE_PATH = "/usr/bin/docker"
	YAML_EXT                       = ".yaml"
	INDEX_FILE_NAME_WITHOUT_EXT    = "index"
	MANIFEST_FILE_NAME             = "manifests"
	INDEX_FILE_NAME_WITH_EXT       = INDEX_FILE_NAME_WITHOUT_EXT + YAML_EXT
	PROCESS_STATUS_RUNNING         = "Active"
	PROCESS_STATUS_STOPPED         = "Inactive"
)

var (
	CONFIG_DIR_PATH   = path.Join(GetHomeDir(), CONFIG_DIR)
	MANIFEST_DIR_PATH = path.Join(CONFIG_DIR_PATH, MANIFEST_FILE_NAME)
	INDEX_FILE_PATH   = path.Join(CONFIG_DIR_PATH, INDEX_FILE_NAME_WITH_EXT)
)

// GetDockerExecutablePath returns executable docker path
func GetDockerExecutablePath() string {
	cmdPath, err := exec.LookPath("docker")
	if err != nil {
		Terminate(fmt.Sprintf("docker doesn't exist probabbly: %s:", err.Error()))
		return DEFAULT_DOCKER_EXECUTABLE_PATH
	}

	if cmdPath, err = filepath.Abs(cmdPath); err != nil {
		Terminate(err.Error())
		return DEFAULT_DOCKER_EXECUTABLE_PATH
	}

	return cmdPath
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
	fmt.Println("err:", str)
	os.Exit(1)
}

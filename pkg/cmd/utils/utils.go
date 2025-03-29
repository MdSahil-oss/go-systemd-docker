package utils

import (
	"go-systemd-docker/pkg/utils"

	"github.com/kardianos/service"
	"github.com/manifoldco/promptui"
)

var logger service.Logger

func PromtForConfirmation(str string) {
	prompt := promptui.Prompt{
		Label:     str,
		IsConfirm: true, // Ensure it's a confirmation prompt
	}

	_, err := prompt.Run()
	if err != nil {
		utils.TerminateWithOutput("Confirmation cancelled.")
	}
}

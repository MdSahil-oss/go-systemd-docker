package utils

import (
	"fmt"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/kardianos/service"
	"github.com/manifoldco/promptui"
)

var logger service.Logger

type CreateProgram struct{}

func (p *CreateProgram) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	// go p.run()
	fmt.Println("Initiated by start function")
	return nil
}

func (p *CreateProgram) run() {
	fmt.Println("this is running inside run function initiated by start function")
	// Do work here
}

func (p *CreateProgram) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	fmt.Println("Stopped running")
	return nil
}

func PromtForConfirmation(str string) {
	prompt := promptui.Prompt{
		Label:     str,
		IsConfirm: true, // Ensure it's a confirmation prompt
	}

	_, err := prompt.Run()
	if err != nil {
		utils.Terminate("Confirmation cancelled.")
	}
}

func GetSystemDProcess(instanceName string) (service.Service, error) {
	var s service.Service
	svcConfig, err := system.GetService(instanceName)
	if err != nil {
		return s, err
	}

	prg := &CreateProgram{}
	s, err = service.New(prg, svcConfig)

	return s, err
}

func GetSystemDProcesses() ([]service.Service, error) {
	var svcs []service.Service

	ss, err := system.ListServices()
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

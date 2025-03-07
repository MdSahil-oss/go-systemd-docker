package cmd

import (
	"fmt"
	"go-systemd-docker/system"

	"github.com/kardianos/service"
)

var logger service.Logger

type createProgram struct{}

func (p *createProgram) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	// go p.run()
	fmt.Println("Initiated by start function")
	return nil
}

func (p *createProgram) run() {
	fmt.Println("this is running inside run function initiated by start function")
	// Do work here
}

func (p *createProgram) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	fmt.Println("Stopped running")
	return nil
}

func GetSystemDProcess(instanceName string) (service.Service, error) {
	var s service.Service
	svcConfig, err := system.GetService(instanceName)
	if err != nil {
		return s, err
	}

	prg := &createProgram{}
	s, err = service.New(prg, svcConfig)

	return s, err

}

package cmd

import (
	"fmt"

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

package cmd

import (
	"fmt"
	"go-systemd-docker/system"
	"go-systemd-docker/utils"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
	"go.rtnl.ai/x/randstr"
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

var createCmd = &cobra.Command{
	Use:     "create [flags]",
	Aliases: []string{"add"},
	Short:   "Register container as Systemd process.",
	Long: `This command registers container as Systemd process.
	e.g. sysd create container-image-name`,
	Args: cobra.RangeArgs(1, 2),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 && len(*flgs.namePersistentFlag) > 0 {
			utils.Terminate("please provide either args[1] or --name not both")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd create")
		// args[0] -> container-image
		imageName := args[0]

		var instanceName string = *flgs.namePersistentFlag
		if len(args) > 1 {
			instanceName = args[1]
		}

		if len(instanceName) == 0 {
			// Assign a random name to `instanceName`.
			instanceName = randstr.Word(8)
		}

		// Do followings:
		// Find a way to start given containerImage (args[0]) as SystemD process.
		// Make `svcConfig` stateful

		// Create -> Check if the instance already exist
		// If not then Calls NewSystem() and create one follows by CreateService

		// *****Learning here*****
		if system.IsServiceExist(instanceName) {
			utils.Terminate(fmt.Sprintf("systemd service already exist with %s", instanceName))
		}

		sysConfig := system.NewSystem(
			system.WithName(instanceName),
			system.WithDisplayName(instanceName),
			system.WithDescription(fmt.Sprintf("Runs %v as %v", instanceName, imageName)),
			system.WithExecutable(""),
			system.WithArguments([]string{}),
			// EnvVars:
		)

		svcConfig, err := system.CreateService(sysConfig)
		if err != nil {
			utils.Terminate(err.Error())
		}

		// prg := &createProgram{}
		// s, err := service.New(prg, svcConfig)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// logger, err = s.Logger(nil)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// // err = s.Run()

		// if err = s.Install(); err != nil {
		// 	logger.Error(err)
		// }

		// Below will be used within `start` after making `svcConfig` stateful
		// if err = s.Start(); err != nil {
		// 	logger.Error(err)
		// }

	},
}

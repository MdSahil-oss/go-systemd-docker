package cmd

import (
	"fmt"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
	"go.rtnl.ai/x/randstr"
)

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
		imageName := args[0]

		var instanceName string = *flgs.namePersistentFlag
		if len(args) > 1 {
			instanceName = args[1]
		}

		if len(instanceName) == 0 {
			// Assign a random name to `instanceName`.
			instanceName = randstr.Word(8)
			AreYouAure(fmt.Sprintf(`Are you sure you want to run '%s' image as systemd process A random name '%s' will be assigned to systemd instance.`, imageName, instanceName))
		}

		if system.IsServiceExist(instanceName) {
			utils.Terminate(fmt.Sprintf("systemd service already exist with %s", instanceName))
		}

		sysConfig := system.NewSystem(
			system.WithName(instanceName),
			system.WithDisplayName(instanceName),
			system.WithDescription(fmt.Sprintf("Runs %v as %v", instanceName, imageName)),
			system.WithExecutable(utils.GetDockerExecutablePath()),
			system.WithArguments([]string{
				"run",
				imageName,
			}),
			// Add other docker supported ...args
		)

		svcConfig, err := system.CreateService(sysConfig, imageName)
		if err != nil {
			utils.Terminate(err.Error())
		}

		prg := &createProgram{}
		s, err := service.New(prg, svcConfig)
		if err != nil {
			// log.Fatal(err)
			utils.Terminate(err.Error())
		}

		logger, err = s.Logger(nil)
		if err != nil {
			// log.Fatal(err)
			utils.Terminate(err.Error())
		}

		// err = s.Run()

		if err = s.Install(); err != nil {
			logger.Error(err)
			utils.Terminate(err.Error())
		}
	},
}

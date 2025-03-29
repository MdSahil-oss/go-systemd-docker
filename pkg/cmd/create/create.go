package create

import (
	"fmt"
	"go-systemd-docker/pkg/system"

	"go-systemd-docker/pkg/utils"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

type FlagsType struct {
	name       *string
	domainName *string
	entrypoint *string
	expose     *[]string
	publish    *[]string
	env        *[]string
}

var CreateCmd *cobra.Command
var flags FlagsType

func init() {
	CreateCmd = &cobra.Command{
		Use:     "create [flags]",
		Aliases: []string{"add"},
		Short:   "Register container as Systemd process.",
		Long: `This command registers container as Systemd process.
		e.g. sysd create nginx
			sysd create nginx sample
			sysd create nginx:latest sample`,
		Args: cobra.RangeArgs(1, 2),
		PreRun: func(command *cobra.Command, args []string) {
			if len(args) > 1 && len(*flags.name) > 0 {
				utils.TerminateWithError("please provide either args[1] or --name not both")
			}
		},
		Run: func(command *cobra.Command, args []string) {
			imageName := args[0]
			var instanceName string = *flags.name
			if len(args) > 1 {
				instanceName = args[1]
			}

			if err := validateImage(command, imageName, instanceName); err != nil {
				utils.TerminateWithError(err.Error())
			}

			if system.IsServiceExist(instanceName) {
				utils.TerminateWithError(fmt.Sprintf("systemd service already exist with %s", instanceName))
			}

			// serializes all flags with respective values.
			dockerFlags := dockerFlagsCollector()
			sysArguments := []string{
				"run",
				"--name",
				instanceName,
			}
			sysArguments = append(sysArguments, dockerFlags...)
			sysArguments = append(sysArguments, imageName)
			sysConfig := system.NewSystem(
				system.WithName(instanceName),
				system.WithDisplayName(instanceName),
				system.WithDescription(fmt.Sprintf("Runs %v as %v", instanceName, imageName)),
				system.WithExecutable(utils.GetDockerExecutablePath()),
				system.WithArguments(sysArguments),
			)

			svcConfig, err := system.CreateService(sysConfig, imageName)
			if err != nil {
				utils.TerminateWithError(err.Error())
			}

			prg := &system.CreateProgram{}
			s, err := service.New(prg, svcConfig)
			if err != nil {
				// log.Fatal(err)
				utils.TerminateWithError(err.Error())
			}

			// logger, err = s.Logger(nil)
			// if err != nil {
			// 	// log.Fatal(err)
			// 	utils.Terminate(err.Error())
			// }

			// err = s.Run()

			if err = s.Install(); err != nil {
				// logger.Error(err)
				utils.TerminateWithError(err.Error())
			}
		},
	}

	flags = FlagsType{
		name:       CreateCmd.Flags().StringP("name", "n", "", "name of the creating instance"),
		domainName: CreateCmd.Flags().StringP("domainname", "d", "", "Container NIS domain name"),
		entrypoint: CreateCmd.Flags().String("entrypoint", "", "Overwrite the default ENTRYPOINT of the image"),
		expose:     CreateCmd.Flags().StringSliceP("expose", "x", []string{}, "Expose a port or a range of ports"),
		publish:    CreateCmd.Flags().StringSliceP("publish", "p", []string{}, "Publish a container's port(s) to the host"),
		env:        CreateCmd.Flags().StringSliceP("env", "e", []string{}, "Set environment variables"),
	}
}

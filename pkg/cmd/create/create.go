package create

import (
	"fmt"
	cmdUtils "go-systemd-docker/pkg/cmd/utils"
	"go-systemd-docker/pkg/system"
	"os/exec"

	"go-systemd-docker/pkg/utils"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
	"go.rtnl.ai/x/randstr"
)

type FlagsType struct {
	name *string
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

			// **Validate the given image if exist by pulling the image locally.**

			// Checks image's local availablity.
			dImagesCmd := exec.Command("docker", "images", "--format=json", imageName)
			dImagesOutput, err := dImagesCmd.CombinedOutput()
			if err != nil {
				utils.TerminateWithError(fmt.Sprintf("docker search: %s\n%s", err.Error(), dImagesOutput))
			}

			if len(dImagesOutput) == 0 {
				// Checks image's registry availablity.
				dSearchCmd := exec.Command("docker", "search", "--format=json", imageName)
				dSearchOutput, err := dSearchCmd.CombinedOutput()
				if err != nil {
					utils.TerminateWithError(fmt.Sprintf("docker search: %s\n%s", err.Error(), dSearchOutput))
				}

				if len(dSearchOutput) == 0 {
					utils.TerminateWithError(fmt.Sprintf("no image found with %s name", imageName))
				}

				// Checks if image is pullable.
				dPullCmd := exec.Command("docker", "pull", imageName)
				dPullOutput, err := dPullCmd.CombinedOutput()
				if err != nil {
					utils.TerminateWithError(fmt.Sprintf("docker pull: %s\n%s", err.Error(), dPullOutput))
				}

				if len(instanceName) == 0 {
					// Assign a random name to `instanceName`.
					instanceName = randstr.Word(8)

					isNotInteractive, err := command.Flags().GetBool("not-interactive")
					if err != nil {
						utils.TerminateWithError(err.Error())
					}

					if !isNotInteractive {
						cmdUtils.PromtForConfirmation(
							fmt.Sprintf(`Are you sure you want to run '%s' image as systemd process A random name '%s' will be assigned to systemd instance.`, imageName, instanceName),
						)
					}
				}
			}

			if system.IsServiceExist(instanceName) {
				utils.TerminateWithError(fmt.Sprintf("systemd service already exist with %s", instanceName))
			}

			sysConfig := system.NewSystem(
				system.WithName(instanceName),
				system.WithDisplayName(instanceName),
				system.WithDescription(fmt.Sprintf("Runs %v as %v", instanceName, imageName)),
				system.WithExecutable(utils.GetDockerExecutablePath()),
				system.WithArguments([]string{
					"run",
					"--name",
					instanceName,
					imageName,
				}),
				// Add other docker supported ...args
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
		name: CreateCmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}
}

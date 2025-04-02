package remove

import (
	"fmt"
	deleteCmd "go-systemd-docker/pkg/cmd/delete"
	stopCmd "go-systemd-docker/pkg/cmd/stop"
	cmdUtils "go-systemd-docker/pkg/cmd/utils"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"
	"io"
	"os/exec"

	"github.com/spf13/cobra"
)

type Flags struct {
}

type remove struct {
	Cmd   *cobra.Command
	Flags Flags
}

func New() *remove {
	remove := &remove{}

	remove.Cmd = &cobra.Command{
		Use:     "remove [flags]",
		Short:   "Remove avaiable images used by systemd process.",
		Aliases: []string{"rm"},
		Long: `This command remove avaiable images used by systemd process..
		e.g. sysd docker rm nginx`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Check if the given image actually being utilized by systemd
			// If so, then first remove all the running process (promt user for confirmation)
			// Then remove the images using docker executable
			cmdUtils.PromtForConfirmation("Removing given image will remove all the running instances, Do you want to continue?")

			for _, element := range args {
				// Validate
				svcs, err := system.ListServices()
				if err != nil {
					utils.TerminateWithError(err.Error())
				}

				if len(svcs) == 0 {
					utils.TerminateWithError(fmt.Sprintf("no image found with %s name", element))
				}

				isImageExist := false
				// Checks the services running on the system using the given image then stop & remove them.
				for _, svc := range svcs {
					if svc.Image == element {
						isImageExist = true
						// Checks if the service is running.
						runningSVC, err := system.ListRunningService(svc.Name)
						if err != nil {
							deleteCmd.New().Cmd.Run(cmd, []string{runningSVC.Name})
							continue
						}

						// Stops && removes.
						stopCmd.New().Cmd.Run(cmd, []string{runningSVC.Name})
						deleteCmd.New().Cmd.Run(cmd, []string{runningSVC.Name})
					}
				}

				// Now remove the image if nothing else is running on the system using the same image.
				if isImageExist {
					dps := exec.Command("docker", "ps", "-a")
					dg := exec.Command("grep", "-i", element)
					rDps, wDps := io.Pipe()

					dps.Stdout = wDps
					dg.Stdin = rDps

					if err := dps.Start(); err != nil {
						utils.TerminateWithError(fmt.Sprintf("couldn't execute 'docker ps -a' : %s", err.Error()))
					}

					go func() {
						dps.Wait()
						wDps.Close()
					}()

					dgOutput, err := dg.CombinedOutput()
					if err != nil {
						utils.TerminateWithError(fmt.Sprintf("couldn't execute 'grep -i' : %s\n%s", err.Error(), string(dgOutput)))
					}

					if len(string(dgOutput)) == 0 {
						drmi := exec.Command("docker", "rmi", "-f", element)
						if err := drmi.Run(); err != nil {
							utils.TerminateWithError(err.Error())
						}
					}
				} else {
					utils.TerminateWithError(fmt.Sprintf("image %s is not utilized by this tool", element))
				}
			}
		},
	}

	return remove
}

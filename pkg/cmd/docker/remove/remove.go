package remove

import (
	"fmt"
	cmdDelete "go-systemd-docker/pkg/cmd/delete"
	"go-systemd-docker/pkg/cmd/stop"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"
	"os/exec"

	"github.com/spf13/cobra"
)

var RemoveCmd *cobra.Command

func init() {
	RemoveCmd = &cobra.Command{
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

			// Validate
			svcs, err := system.ListServices()
			if err != nil {
				utils.Terminate(err.Error())
			}

			if len(svcs) == 0 {
				utils.Terminate(fmt.Sprintf("no image found with %s name", args[0]))
			}

			isImageExist := false
			// Checks the services running on the system using the given image then stop & remove them.
			for _, svc := range svcs {
				if svc.Image == args[0] {
					isImageExist = true
					// Checks if the service is running.
					runningSVC, err := system.ListRunningService(svc.Name)
					if err != nil {
						cmdDelete.DeleteCmd.Run(cmd, []string{runningSVC.Name})
						continue
					}

					// Stops && removes.
					stop.StopCmd.Run(cmd, []string{runningSVC.Name})
					cmdDelete.DeleteCmd.Run(cmd, []string{runningSVC.Name})
				}
			}

			// Now remove the image if nothing else is running on the system using the same image.
			if isImageExist {
				dps := exec.Command("bash", "-c", "docker ps | grep -i", args[0])
				dpsOutput, err := dps.CombinedOutput()
				if err != nil {
					utils.Terminate(err.Error())
				}

				if len(string(dpsOutput)) == 0 {
					drmi := exec.Command("bash", "-c", "docker rmi -f", args[0])
					if _, err := drmi.CombinedOutput(); err != nil {
						utils.Terminate(err.Error())
					}
				}
			} else {
				utils.Terminate(fmt.Sprintf("image %s is not utilized by this tool", args[0]))
			}
		},
	}
}

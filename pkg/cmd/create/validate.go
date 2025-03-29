package create

import (
	"fmt"
	cmdUtils "go-systemd-docker/pkg/cmd/utils"
	"os/exec"

	"github.com/spf13/cobra"
	"go.rtnl.ai/x/randstr"
)

// Validate the given image if exist by pulling the image locally
func validateImage(cmd *cobra.Command, imageName, instanceName string) error {
	// Checks image's local availablity.
	dImagesCmd := exec.Command("docker", "images", "--format=json", imageName)
	dImagesOutput, err := dImagesCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker search: %s\n%s", err.Error(), dImagesOutput)
	}

	if len(dImagesOutput) == 0 {
		// Checks image's registry availablity.
		dSearchCmd := exec.Command("docker", "search", "--format=json", imageName)
		dSearchOutput, err := dSearchCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("docker search: %s\n%s", err.Error(), dSearchOutput)
		}

		if len(dSearchOutput) == 0 {
			return fmt.Errorf("no image found with %s name", imageName)
		}

		// Checks if image is pullable.
		dPullCmd := exec.Command("docker", "pull", imageName)
		dPullOutput, err := dPullCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("docker pull: %s\n%s", err.Error(), dPullOutput)
		}

		if len(instanceName) == 0 {
			// Assign a random name to `instanceName`.
			instanceName = randstr.Word(8)

			isNotInteractive, err := cmd.Flags().GetBool("not-interactive")
			if err != nil {
				return err
			}

			if !isNotInteractive {
				cmdUtils.PromtForConfirmation(
					fmt.Sprintf(`Are you sure you want to run '%s' image as systemd process A random name '%s' will be assigned to systemd instance.`, imageName, instanceName),
				)
			}
		}
	}

	return nil
}

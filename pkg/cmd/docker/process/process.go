package process

import (
	dockerUtils "go-systemd-docker/pkg/cmd/docker/utils"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

var ProcessCmd *cobra.Command

func init() {
	ProcessCmd = &cobra.Command{
		Use:     "process [flags]",
		Short:   "List running systemd process images.",
		Aliases: []string{"ps"},
		Long: `This command List running systemd process images.
		e.g. sysd docker ps`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				utils.TerminateWithError("no argument was expected")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			isvcs, err := system.ListRunningServices()
			if err != nil {
				utils.TerminateWithError(err.Error())
			}

			if len(isvcs) == 0 {
				utils.TerminateWithError("no running image found")
			}

			dockerUtils.PrintImagesFromIndexService(isvcs)
		},
	}
}

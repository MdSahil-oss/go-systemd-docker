package list

import (
	dockerUtils "go-systemd-docker/pkg/cmd/docker/utils"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

var ListCmd *cobra.Command

func init() {
	ListCmd = &cobra.Command{
		Use:     "list [flags]",
		Short:   "List containers used by systemd proces(es)",
		Aliases: []string{"ls"},
		Long: `This command List containers used by systemd proces(es).
		e.g. sysd docker ls`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				utils.TerminateWithError("no argument was expected")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			svcs, err := system.ListServices()
			if err != nil {
				utils.TerminateWithError(err.Error())
			}

			if len(svcs) == 0 {
				utils.TerminateWithError("no image found")
			}

			dockerUtils.PrintImagesFromIndexService(svcs)
		},
	}
}

package docker

import (
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

var processCmd = &cobra.Command{
	Use:     "process [flags]",
	Short:   "List running systemd process (container-image).",
	Aliases: []string{"ps"},
	Long: `This command List running systemd process (container-image).
	e.g. sysd ps`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			utils.Terminate("no argument was expected")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

package remove

import (
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
		},
	}
}

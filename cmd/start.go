package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [flags]",
	Short: "Continue running stopped systemd process (container-image).",
	Long: `This command continue running stopped systemd process (container-image).
	e.g. sysd start registered-systemd-instance-name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd start")

		if *flgs.namePersistentFlag != "" {
			fmt.Println("Provided name:", *flgs.namePersistentFlag)
		}
	},
}

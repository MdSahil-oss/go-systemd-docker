package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop [flags]",
	Short: "Stop running systemd process (container-image).",
	Long: `This command stop running systemd process (container-image).
	e.g. sysd stop registered-systemd-instance-name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd stop")

		if *flgs.namePersistentFlag != "" {
			fmt.Println("Provided name:", *flgs.namePersistentFlag)
		}
	},
}

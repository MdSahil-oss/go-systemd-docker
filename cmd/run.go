package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [flags]",
	Short: "Initiate registered container running as Systemd process.",
	Long: `This command creates an instance of registered container as Systemd process.
	e.g. sysd run registered-container-name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd run")

		if *flgs.namePersistentFlag != "" {
			fmt.Println("Provided name:", *flgs.namePersistentFlag)
		}
	},
}

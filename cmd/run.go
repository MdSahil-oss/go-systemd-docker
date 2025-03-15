package cmd

import (
	"go-systemd-docker/utils"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [flags]",
	Short: "Initiate registered container running as Systemd process.",
	Long: `This command creates an instance of registered container as Systemd process.
	e.g. sysd run registered-container-name`,
	Args: cobra.RangeArgs(1, 2),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 && len(*flgs.namePersistentFlag) > 0 {
			utils.Terminate("please provide either args[1] or --name not both")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		createCmd.Run(cmd, args)

		var instanceName string = *flgs.namePersistentFlag
		if len(args) > 1 {
			instanceName = args[1]
		}

		startCmd.Run(cmd, []string{instanceName})
	},
}

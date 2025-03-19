package cmd

import (
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start [flags]",
	Short: "Continue running stopped systemd process (container-image).",
	Long: `This command continue running stopped systemd process (container-image).
	e.g. sysd start registered-systemd-instance-name`,
	Args: cobra.RangeArgs(0, 1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(*flgs.namePersistentFlag) > 0 {
			utils.Terminate("please provide either args[0] or --name not both")
		}

		if len(args) == 0 && len(*flgs.namePersistentFlag) == 0 {
			utils.Terminate("please provide either args[0] or --name")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var instanceName string = *flgs.namePersistentFlag
		if len(args) > 0 {
			instanceName = args[0]
		}

		svc, err := GetSystemDProcess(instanceName)
		if err != nil {
			utils.Terminate(err.Error())
		}

		logger, err = svc.Logger(nil)
		if err != nil {
			utils.Terminate(err.Error())
			// log.Fatal(err)
		}

		if err := svc.Start(); err != nil {
			logger.Error(err)
			utils.Terminate(err.Error())
		}

		// if err = svc.Run(); err != nil {
		// 	utils.Terminate(err.Error())
		// }
	},
}

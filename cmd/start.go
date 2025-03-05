package cmd

import (
	"go-systemd-docker/system"
	"go-systemd-docker/utils"

	"github.com/kardianos/service"
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

		svcConfig, err := system.GetService(instanceName)
		if err != nil {
			utils.Terminate(err.Error())
		}

		prg := &createProgram{}
		s, err := service.New(prg, svcConfig)
		if err != nil {
			utils.Terminate(err.Error())
			// log.Fatal(err)
		}

		logger, err = s.Logger(nil)
		if err != nil {
			utils.Terminate(err.Error())
			// log.Fatal(err)
		}

		if err := s.Start(); err != nil {
			logger.Error(err)
			utils.Terminate(err.Error())
		}
	},
}

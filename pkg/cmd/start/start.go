package start

import (
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

type Flags struct {
	name *string
}

type Start struct {
	Cmd   *cobra.Command
	Flags Flags
}

func New() *Start {
	start := &Start{}

	start.Cmd = &cobra.Command{
		Use:   "start [flags]",
		Short: "Continue running stopped systemd process (container-image).",
		Long: `This command continue running stopped systemd process (container-image).
		e.g. sysd start registered-systemd-instance-name`,
		Args: cobra.RangeArgs(0, 1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && len(*start.Flags.name) > 0 {
				utils.TerminateWithError("please provide either args[0] or --name not both")
			}

			if len(args) == 0 && len(*start.Flags.name) == 0 {
				utils.TerminateWithError("please provide either args[0] or --name")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var instanceName string = *start.Flags.name
			if len(args) > 0 {
				instanceName = args[0]
			}

			svc, err := system.GetSystemDProcess(instanceName)
			if err != nil {
				utils.TerminateWithError(err.Error())
			}

			// logger, err = svc.Logger(nil)
			// if err != nil {
			// 	utils.TerminateWithError(err.Error())
			// 	// log.Fatal(err)
			// }

			if err := svc.Start(); err != nil {
				// logger.Error(err)
				utils.TerminateWithError(err.Error())
			}

			// if err = svc.Run(); err != nil {
			// 	utils.TerminateWithError(err.Error())
			// }
		},
	}

	start.Flags = Flags{
		name: start.Cmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}

	return start
}

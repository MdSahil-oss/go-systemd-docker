package stop

import (
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

type Flags struct {
	name *string
}

type Stop struct {
	Cmd   *cobra.Command
	Flags Flags
}

func New() *Stop {
	stop := &Stop{}

	stop.Cmd = &cobra.Command{
		Use:   "stop [flags]",
		Short: "Stop running systemd process (container-image).",
		Long: `This command stop running systemd process (container-image).
		e.g. sysd stop registered-systemd-instance-name`,
		Args: cobra.RangeArgs(0, 1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && len(*stop.Flags.name) > 0 {
				utils.TerminateWithError("please provide either args[0] or --name not both")
			}

			if len(args) == 0 && len(*stop.Flags.name) == 0 {
				utils.TerminateWithError("please provide either args[0] or --name")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var instanceName string = *stop.Flags.name
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

			if err := svc.Stop(); err != nil {
				// logger.Error(err)
				utils.TerminateWithError(err.Error())
			}
		},
	}

	stop.Flags = Flags{
		name: stop.Cmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}

	return stop
}

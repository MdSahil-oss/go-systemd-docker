package stop

import (
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

type FlagsType struct {
	name *string
}

var StopCmd *cobra.Command
var flags FlagsType

func init() {
	StopCmd = &cobra.Command{
		Use:   "stop [flags]",
		Short: "Stop running systemd process (container-image).",
		Long: `This command stop running systemd process (container-image).
		e.g. sysd stop registered-systemd-instance-name`,
		Args: cobra.RangeArgs(0, 1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && len(*flags.name) > 0 {
				utils.TerminateWithError("please provide either args[0] or --name not both")
			}

			if len(args) == 0 && len(*flags.name) == 0 {
				utils.TerminateWithError("please provide either args[0] or --name")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var instanceName string = *flags.name
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

	flags = FlagsType{
		name: StopCmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}
}

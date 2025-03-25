package run

import (
	"go-systemd-docker/pkg/cmd/create"
	"go-systemd-docker/pkg/cmd/start"
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

type FlagsType struct {
	name *string
}

var RunCmd *cobra.Command
var flags FlagsType

func init() {
	RunCmd = &cobra.Command{
		Use:   "run [flags]",
		Short: "Initiate registered container running as Systemd process.",
		Long: `This command creates an instance of registered container as Systemd process.
		e.g. sysd run registered-container-name`,
		Args: cobra.RangeArgs(1, 2),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 && len(*flags.name) > 0 {
				utils.Terminate("please provide either args[1] or --name not both")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var instanceName string = *flags.name
			if len(args) > 1 {
				instanceName = args[1]
			}

			create.CreateCmd.Run(cmd, []string{args[0], instanceName})
			start.StartCmd.Run(cmd, []string{instanceName})
		},
	}

	flags = FlagsType{
		name: RunCmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}
}

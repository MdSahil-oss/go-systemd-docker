package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"go.rtnl.ai/x/randstr"
)

var createCmd = &cobra.Command{
	Use:     "create [flags]",
	Aliases: []string{"add"},
	Short:   "Register container as Systemd process.",
	Long: `This command registers container as Systemd process.
	e.g. sysd create container-image-name`,
	Args: cobra.RangeArgs(1, 2),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 1 && len(*flgs.namePersistentFlag) > 0 {
			fmt.Println("please provide either args[1] or --name not both")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd create")
		// args[0] -> container-image

		var instanceName string = *flgs.namePersistentFlag
		if len(args) > 1 {
			instanceName = args[1]
		}

		if len(instanceName) == 0 {
			// Assign a random name to `instanceName`.
			instanceName = randstr.Word(8)
		}

		// Do followings:
		// Find a way to start given containerImage (args[0]) as SystemD process.
	},
}

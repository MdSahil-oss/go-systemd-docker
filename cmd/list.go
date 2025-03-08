package cmd

import (
	"go-systemd-docker/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List registered systemd process (container-image).",
	Aliases: []string{"ls"},
	Long: `This command List registered systemd process (container-image).
	e.g. sysd ls`,
	Args: cobra.RangeArgs(0, 1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(*flgs.namePersistentFlag) > 0 {
			utils.Terminate("please provide either args[0] or --name not both")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// var instanceName string = *flgs.namePersistentFlag
		// if len(args) > 0 {
		// 	instanceName = args[0]
		// }

	},
}

// name: "index"
// services: [
// {
// 	name: "sample",
// 	path: "./manifests/sample.yaml"
// },
// {
// 	name: "sample-1",
// 	path: "./manifests/sample-1.yaml"
// }
// ]

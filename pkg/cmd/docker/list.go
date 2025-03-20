package docker

import (
	"fmt"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List containers used by systemd proces(es)",
	Aliases: []string{"ls"},
	Long: `This command List containers used by systemd proces(es).
	e.g. sysd docker ls`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			utils.Terminate("no argument was expected")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		svcs, err := system.ListServices()
		if err != nil {
			utils.Terminate(err.Error())
		}

		// Unifying the images
		set := make(map[string]any)
		for _, svc := range svcs {
			set[svc.Image] = struct{}{}
		}

		if len(set) == 0 {
			utils.Terminate("no image found")
		}
		for k, _ := range set {
			fmt.Println(k)
		}
	},
}

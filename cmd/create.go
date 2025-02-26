package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create [flags]",
	Aliases: []string{"add"},
	Short:   "Register container as Systemd process.",
	Long: `This command registers container as Systemd process.
	e.g. sysd create container-image-name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd create")
	},
}

func init() {

}

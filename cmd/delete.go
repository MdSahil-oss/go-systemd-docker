package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [flags]",
	Aliases: []string{"rm", "remove"},
	Short:   "Deregister  already register container as Systemd process.",
	Long: `This command deregisters container as Systemd process.
	e.g. sysd delete container-image-name`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd delete...")
	},
}

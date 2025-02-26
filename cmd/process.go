package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var processCmd = &cobra.Command{
	Use:     "process [flags]",
	Short:   "List running systemd process (container-image).",
	Aliases: []string{"ps"},
	Long: `This command List running systemd process (container-image).
	e.g. sysd ps`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd ps")
	},
}

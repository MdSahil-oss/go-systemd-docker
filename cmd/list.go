package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List registered systemd process (container-image).",
	Aliases: []string{"ls"},
	Long: `This command List registered systemd process (container-image).
	e.g. sysd ls`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sysd ls")

		if *flgs.namePersistentFlag != "" {
			fmt.Println("Provided name:", *flgs.namePersistentFlag)
		}
	},
}

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "sysd [flags] [cmd]",
	Aliases: []string{"systemd-docker"},
	Short:   "Sysd manages Docker container as Systemd processes",
	Long: `A CLI utility which manages Docker images as SystemD processes.
	e.g. sysd stands for systemd-docker`,
	Version: "1.0.0",
}

type Flags struct {
	namePersistentFlag *string
	forceFlag          *bool
	allFlag            *bool
}

var flgs = Flags{}

// Using init() to register flags with sub-cmds
func init() {
	// persistent Flags for all the sub-cmds.
	flgs.namePersistentFlag = rootCmd.PersistentFlags().StringP("name", "n", "", "provides name to instance")

	// flags for delete (or rm) sub-cmd
	flgs.forceFlag = deleteCmd.Flags().BoolP("force", "f", false, "force delete packages")

	// flags for process (or ps) sub-cmd
	flgs.allFlag = processCmd.Flags().BoolP("all", "a", false, "select all packages")
}

func Execute() {
	// v.SetDefault("random", "It should not have been random")
	// v := viper.GetViper()
	// fmt.Println("random:", v.GetString("random"))

	// name := rootCmd.PersistentFlags().StringP("name", "n", "", "used for providing name")
	// v := viper.New()
	// v.BindPFlag()

	// registers cmds.
	rootCmd.AddCommand(
		createCmd,
		deleteCmd,
		runCmd,
		startCmd,
		stopCmd,
		listCmd,
		processCmd,
		showCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

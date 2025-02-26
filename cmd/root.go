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

func init() {
	fmt.Println("Init running for root")
}

func Execute() {
	// registers cmds.
	rootCmd.AddCommand(
		createCmd,
		deleteCmd,
		runCmd,
		startCmd,
		stopCmd,
		listCmd,
		processCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

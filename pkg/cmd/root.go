package cmd

import (
	"fmt"
	"go-systemd-docker/pkg/cmd/docker"
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
	namePersistentFlag           *string
	notInteractivePersistentFlag *bool
	forceFlag                    *bool
	allFlag                      *bool
}

var flgs = Flags{}

// Using init() to register flags with sub-cmds
func init() {
	// persistent Flags for all the sub-cmds.
	flgs.namePersistentFlag = rootCmd.PersistentFlags().StringP("name", "n", "", "provides name to instance")
	flgs.notInteractivePersistentFlag = rootCmd.PersistentFlags().BoolP("not-interactive", "t", false, "enables non-interactive mode")
	flgs.allFlag = rootCmd.PersistentFlags().BoolP("all", "a", false, "select all packages")

	// flags for delete (or rm) sub-cmd
	flgs.forceFlag = deleteCmd.Flags().BoolP("force", "f", false, "force delete packages")

	// flags for process (or ps) sub-cmd
}

func Execute() {
	// Registers groups
	rootCmd.AddGroup(docker.DockerGrp)

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
		docker.DockerCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

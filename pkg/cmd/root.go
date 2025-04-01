package cmd

import (
	"fmt"
	"go-systemd-docker/pkg/cmd/create"
	"go-systemd-docker/pkg/cmd/delete"
	"go-systemd-docker/pkg/cmd/docker"
	"go-systemd-docker/pkg/cmd/list"
	"go-systemd-docker/pkg/cmd/process"
	"go-systemd-docker/pkg/cmd/run"
	"go-systemd-docker/pkg/cmd/show"
	"go-systemd-docker/pkg/cmd/start"
	"go-systemd-docker/pkg/cmd/stop"
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

type FlagsType struct {
	NotInteractivePersistentFlag *bool
}

var Flags = FlagsType{}

// Using init() to register flags with sub-cmds
func init() {
	// persistent Flags for all the sub-cmds.
	Flags.NotInteractivePersistentFlag = rootCmd.PersistentFlags().BoolP("not-interactive", "t", false, "enables non-interactive mode")
}

func Execute() {
	// Collecting all the cmds with flags as components here.
	dockerComp := docker.New()

	// Registers groups
	rootCmd.AddGroup(dockerComp.Group)

	// registers cmds.
	rootCmd.AddCommand(
		dockerComp.Cmd,
		create.New().Cmd,
		delete.New().Cmd,
		list.New().Cmd,
		run.New().Cmd,
		show.New().Cmd,
		start.New().Cmd,
		stop.New().Cmd,
		process.New().Cmd,
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

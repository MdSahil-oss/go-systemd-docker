package docker

import (
	"go-systemd-docker/pkg/cmd/docker/list"
	"go-systemd-docker/pkg/cmd/docker/process"
	"go-systemd-docker/pkg/cmd/docker/remove"

	"github.com/spf13/cobra"
)

type Flags struct {
	namePersistentFlag *string
	forceFlag          *bool
	allFlag            *bool
}

type Docker struct {
	Cmd   *cobra.Command
	Flags Flags
	Group *cobra.Group
}

func New() *Docker {
	docker := &Docker{}

	docker.Group = &cobra.Group{
		ID:    "docker",
		Title: "Docker sub-cmd for containers management",
	}

	docker.Cmd = &cobra.Command{
		Aliases: []string{"d", "do"},
		GroupID: docker.Group.ID,
		Use:     "docker [flags] [cmd]",
		Short:   "A group of commands to manage docker container",
		Long: `Docker sub-cmd is a group of commands to manage docker container used by Systemd processes.
		e.g. sysd docker ps # To list docker images currently in use
		sysd docker ls # To list docker images manage by sysd
		sysd docker rm # To delete one or more images used by CLI`,
	}

	// persistent Flags for all the sub-cmds.
	docker.Flags.namePersistentFlag = docker.Cmd.PersistentFlags().StringP("name", "n", "", "provides name to instance")
	docker.Flags.allFlag = docker.Cmd.PersistentFlags().BoolP("all", "a", false, "select all packages")

	// flags for process (or ps) sub-cmd
	docker.Cmd.AddCommand(list.New().Cmd)
	docker.Cmd.AddCommand(process.New().Cmd)
	docker.Cmd.AddCommand(remove.New().Cmd)

	return docker
}

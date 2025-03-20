package docker

import "github.com/spf13/cobra"

var DockerGrp = &cobra.Group{
	ID:    "docker",
	Title: "Docker sub-cmd for containers management",
}

var DockerCmd = &cobra.Command{
	Aliases: []string{"d", "do"},
	GroupID: DockerGrp.ID,
	Use:     "docker [flags] [cmd]",
	Short:   "A group of commands to manage docker container",
	Long: `Docker sub-cmd is a group of commands to manage docker container used by Systemd processes.
	e.g. sysd docker ps # To list docker images currently in use
	sysd docker ls # To list docker images manage by sysd
	sysd docker rm # To delete one or more images used by CLI`,
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
	flgs.namePersistentFlag = DockerCmd.PersistentFlags().StringP("name", "n", "", "provides name to instance")
	flgs.allFlag = DockerCmd.PersistentFlags().BoolP("all", "a", false, "select all packages")

	// flags for process (or ps) sub-cmd

	DockerCmd.AddCommand(listCmd)
	// DockerCmd.AddCommand(processCmd)
	// DockerCmd.AddCommand(removeCmd)
}

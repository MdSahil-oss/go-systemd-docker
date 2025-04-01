package run

import (
	"go-systemd-docker/pkg/cmd/create"
	"go-systemd-docker/pkg/cmd/start"
	"go-systemd-docker/pkg/utils"

	"github.com/spf13/cobra"
)

type Flags struct {
	name       *string
	domainName *string
	entrypoint *string
	expose     *[]string
	publish    *[]string
	env        *[]string
}

type Run struct {
	Cmd   *cobra.Command
	Flags Flags
}

func New() *Run {
	run := &Run{}

	run.Cmd = &cobra.Command{
		Use:   "run [flags]",
		Short: "Initiate registered container running as Systemd process.",
		Long: `This command creates an instance of registered container as Systemd process.
		e.g. sysd run registered-container-name`,
		Args: cobra.RangeArgs(1, 2),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 1 && len(*run.Flags.name) > 0 {
				utils.TerminateWithError("please provide either args[1] or --name not both")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var instanceName string = *run.Flags.name
			if len(args) > 1 {
				instanceName = args[1]
			}

			// You should have some kind of a way here to have an instance of `create`
			// so that you could assign values to flags and then call Run() of than instance.
			createComp := create.New()
			create.UpdateFlags(
				&createComp.Flags,
				create.WithDomainName(run.Flags.domainName),
				create.WithEntrypoint(run.Flags.entrypoint),
				create.WithEnv(run.Flags.env),
				create.WithExpose(run.Flags.expose),
				create.WithPublish(run.Flags.publish),
			)
			createComp.Cmd.Run(cmd, []string{args[0], instanceName})
			start.New().Cmd.Run(cmd, []string{instanceName})
		},
	}

	run.Flags = Flags{
		name:       run.Cmd.Flags().StringP("name", "n", "", "name of the creating instance"),
		domainName: run.Cmd.Flags().StringP("domainname", "d", "", "Container NIS domain name"),
		entrypoint: run.Cmd.Flags().String("entrypoint", "", "Overwrite the default ENTRYPOINT of the image"),
		expose:     run.Cmd.Flags().StringSliceP("expose", "x", []string{}, "Expose a port or a range of ports"),
		publish:    run.Cmd.Flags().StringSliceP("publish", "p", []string{}, "Publish a container's port(s) to the host"),
		env:        run.Cmd.Flags().StringSliceP("env", "e", []string{}, "Set environment variables"),
	}

	return run
}

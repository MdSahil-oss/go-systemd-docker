package process

import (
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"
	"os"

	"github.com/kataras/tableprinter"
	"github.com/kataras/tablewriter"
	"github.com/spf13/cobra"
)

type Flags struct {
	name *string
	all  *bool
}

type Process struct {
	Cmd   *cobra.Command
	Flags Flags
}

func New() *Process {
	ps := &Process{}

	ps.Cmd = &cobra.Command{
		Use:     "process [flags]",
		Short:   "List running systemd process (container-image).",
		Aliases: []string{"ps"},
		Long: `This command List running systemd process (container-image).
		e.g. sysd ps`,
		Args: cobra.RangeArgs(0, 1),
		PreRun: func(command *cobra.Command, args []string) {
			if len(args) > 0 && len(*ps.Flags.name) > 0 {
				utils.TerminateWithError("please provide either args[0] or --name not both")
			}
		},
		Run: func(command *cobra.Command, args []string) {
			var instanceName string = *ps.Flags.name
			if len(args) > 0 {
				instanceName = args[0]
			}

			// TablePrinter configuration
			printer := tableprinter.New(os.Stdout)
			printer.HeaderLine = false
			printer.HeaderFgColor = tablewriter.FgGreenColor
			var runningSS []system.IndexService

			if len(instanceName) > 0 {
				is, err := system.ListRunningService(instanceName)
				if err != nil {
					utils.TerminateWithError(err.Error())
				}

				runningSS = append(runningSS, is)
			} else {
				ss, err := system.ListRunningServices()
				if err != nil {
					utils.TerminateWithError(err.Error())
				}

				runningSS = ss
			}

			printer.Print(runningSS)
		},
	}

	ps.Flags = Flags{
		name: ps.Cmd.Flags().StringP("name", "n", "", "name of the creating sysd process"),
		all:  ps.Cmd.Flags().BoolP("all", "a", false, "select all sysd process to list"),
	}

	return ps
}

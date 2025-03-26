package process

import (
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"
	"os"

	"github.com/kataras/tableprinter"
	"github.com/kataras/tablewriter"
	"github.com/spf13/cobra"
)

type FlagsType struct {
	name *string
	all  *bool
}

var ProcessCmd *cobra.Command
var flags FlagsType

func init() {
	ProcessCmd = &cobra.Command{
		Use:     "process [flags]",
		Short:   "List running systemd process (container-image).",
		Aliases: []string{"ps"},
		Long: `This command List running systemd process (container-image).
		e.g. sysd ps`,
		Args: cobra.RangeArgs(0, 1),
		PreRun: func(command *cobra.Command, args []string) {
			if len(args) > 0 && len(*flags.name) > 0 {
				utils.Terminate("please provide either args[0] or --name not both")
			}
		},
		Run: func(command *cobra.Command, args []string) {
			var instanceName string = *flags.name
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
					utils.Terminate(err.Error())
				}

				runningSS = append(runningSS, is)
			} else {
				ss, err := system.ListRunningServices()
				if err != nil {
					utils.Terminate(err.Error())
				}

				runningSS = ss
			}

			printer.Print(runningSS)
		},
	}

	flags = FlagsType{
		name: ProcessCmd.Flags().StringP("name", "n", "", "name of the creating sysd process"),
		all:  ProcessCmd.Flags().BoolP("all", "a", false, "select all sysd process to list"),
	}
}

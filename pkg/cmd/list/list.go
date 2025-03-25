package list

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
}

var ListCmd *cobra.Command
var flags FlagsType

func init() {
	ListCmd = &cobra.Command{
		Use:     "list [flags]",
		Short:   "List registered systemd process (container-image).",
		Aliases: []string{"ls"},
		Long: `This command List registered systemd process (container-image).
		e.g. sysd ls`,
		Args: cobra.RangeArgs(0, 1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && len(*flags.name) > 0 {
				utils.Terminate("please provide either args[0] or --name not both")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var instanceName string = *flags.name
			if len(args) > 0 {
				instanceName = args[0]
			}

			svcs, err := system.ListServices()
			if err != nil {
				utils.Terminate(err.Error())
			}

			printer := tableprinter.New(os.Stdout)
			printer.HeaderLine = false
			printer.HeaderFgColor = tablewriter.FgGreenColor

			if len(instanceName) > 0 {
				for i, svc := range svcs {
					if svc.Name == instanceName {
						printer.Print(
							svcs[i],
						)
					}
				}
			} else {
				printer.Print(svcs)
			}
		},
	}

	flags = FlagsType{
		name: ListCmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}
}

package list

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
}

type List struct {
	Cmd   *cobra.Command
	Flags Flags
}

func New() *List {
	list := &List{}

	list.Cmd = &cobra.Command{
		Use:     "list [flags]",
		Short:   "List registered systemd process (container-image).",
		Aliases: []string{"ls"},
		Long: `This command List registered systemd process (container-image).
		e.g. sysd ls`,
		Args: cobra.RangeArgs(0, 1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && len(*list.Flags.name) > 0 {
				utils.TerminateWithError("please provide either args[0] or --name not both")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var instanceName string = *list.Flags.name
			if len(args) > 0 {
				instanceName = args[0]
			}

			svcs, err := system.ListServices()
			if err != nil {
				utils.TerminateWithError(err.Error())
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

	list.Flags = Flags{
		name: list.Cmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}

	return list
}

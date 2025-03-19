package cmd

import (
	"fmt"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"
	"os"

	"github.com/kardianos/service"
	"github.com/kataras/tableprinter"
	"github.com/kataras/tablewriter"
	"github.com/spf13/cobra"
)

var processCmd = &cobra.Command{
	Use:     "process [flags]",
	Short:   "List running systemd process (container-image).",
	Aliases: []string{"ps"},
	Long: `This command List running systemd process (container-image).
	e.g. sysd ps`,
	Args: cobra.RangeArgs(0, 1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(*flgs.namePersistentFlag) > 0 {
			utils.Terminate("please provide either args[0] or --name not both")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var instanceName string = *flgs.namePersistentFlag
		if len(args) > 0 {
			instanceName = args[0]
		}

		// TablePrinter configuration
		printer := tableprinter.New(os.Stdout)
		printer.HeaderLine = false
		printer.HeaderFgColor = tablewriter.FgGreenColor

		if len(instanceName) > 0 {
			s, err := system.ListService(instanceName)
			if err != nil {
				utils.Terminate(err.Error())
			}

			svc, err := GetSystemDProcess(instanceName)
			if err != nil {
				utils.Terminate(err.Error())
			}

			status, err := svc.Status()
			if err != nil {
				utils.Terminate(err.Error())
			}

			if status == service.StatusRunning {
				s.Status = utils.PROCESS_STATUS_RUNNING
				printer.Print(s)
			}

		} else {
			ss, err := system.ListServices()
			if err != nil {
				utils.Terminate(err.Error())
			}

			var runningSS []system.IndexService
			var errs []error = nil
			for _, s := range ss {
				svc, err := GetSystemDProcess(s.Name)
				if err != nil {
					errs = append(errs, err)
				}

				status, err := svc.Status()
				if err != nil {
					errs = append(errs, err)
				}

				if status == service.StatusRunning {
					s.Status = utils.PROCESS_STATUS_RUNNING
					runningSS = append(runningSS, s)
				} else if *flgs.allFlag && status == service.StatusStopped {
					s.Status = utils.PROCESS_STATUS_STOPPED
					runningSS = append(runningSS, s)
				}
			}
			printer.Print(runningSS)

			if len(errs) > 0 {
				utils.Terminate(fmt.Sprintf("%v", errs))
			}
		}
	},
}

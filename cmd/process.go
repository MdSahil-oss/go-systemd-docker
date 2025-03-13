package cmd

import (
	"fmt"
	"go-systemd-docker/system"
	"go-systemd-docker/utils"
	"os"

	"github.com/kardianos/service"
	"github.com/kataras/tableprinter"
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
	// 	--all | -a : list all stopped and running ones.
	// --name | -n : Name of a particular systemd instance to list
	Run: func(cmd *cobra.Command, args []string) {
		var instanceName string = *flgs.namePersistentFlag
		if len(args) > 0 {
			instanceName = args[0]
		}

		tp := tableprinter.New(os.Stdout)
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
				tp.Print(s.Name)
				tp.Print(s.Status)
			}

		} else {
			ss, err := system.ListServices()
			if err != nil {
				utils.Terminate(err.Error())
			}

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
					tp.Print(s.Name)
					tp.Print(s.Status)
				}

				// if status == service.StatusStopped {
				// 	// fmt.Println(fmt.Sprintf("Service %s is stopped", instanceName))
				// 	s.Status = utils.PROCESS_STATUS_STOPPED
				// 	tp.Print(s.Name)
				// 	tp.Print(s.Status)
				// }
			}

			if len(errs) > 0 {
				utils.Terminate(fmt.Sprintf("%v", errs))
			}
		}
	},
}

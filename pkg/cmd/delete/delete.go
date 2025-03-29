package delete

import (
	"fmt"
	cmdUtils "go-systemd-docker/pkg/cmd/utils"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

type FlagsType struct {
	force *bool
	name  *string
	all   *bool
}

var DeleteCmd *cobra.Command
var flags FlagsType

func init() {
	DeleteCmd = &cobra.Command{
		Use:     "delete [flags]",
		Aliases: []string{"rm", "remove"},
		Short:   "Deregister  already register container as Systemd process.",
		Long: `This command deregisters container as Systemd process.
		e.g. sysd delete container-image-name`,
		Args: cobra.MinimumNArgs(0),
		PreRun: func(command *cobra.Command, args []string) {
			if len(args) > 0 && len(*flags.name) > 0 {
				utils.TerminateWithError("please provide either args[0] or --name not both")
			}

			if !*flags.all &&
				len(args) == 0 &&
				len(*flags.name) == 0 {
				utils.TerminateWithError("please provide either args[0] or --name")
			}
		},
		Run: func(command *cobra.Command, args []string) {
			var errs []error
			var instanceNames = []string{*flags.name}
			if *flags.all {
				cmdUtils.PromtForConfirmation("Are you sure? to delete all the instances")
				instanceNames = nil
				svcs, err := system.ListServices()
				if err != nil {
					utils.TerminateWithError(err.Error())
				}

				for _, svc := range svcs {
					instanceNames = append(instanceNames, svc.Name)
				}

			} else if len(args) > 0 {
				instanceNames = args
			}

			for _, instanceName := range instanceNames {
				svc, err := system.GetSystemDProcess(instanceName)
				if err != nil {
					errs = append(errs, err)
					continue
				}

				svcStatus, err := svc.Status()
				if err != nil {
					errs = append(errs, err)
					continue
				}

				if !*flags.force && svcStatus == service.StatusRunning {
					errs = append(errs, fmt.Errorf("service %s is running, please force stop it using '-f'", instanceName))
					continue
				}

				// logger, err = svc.Logger(nil)
				// if err != nil {
				// 	utils.Terminate(err.Error())
				// 	// log.Fatal(err)
				// }

				if svcStatus == service.StatusRunning {
					if err := svc.Stop(); err != nil {
						// logger.Error(err)
						errs = append(errs, err)
						continue
					}
				}

				if err := svc.Uninstall(); err != nil {
					// logger.Error(err)
					errs = append(errs, err)
					continue
				}

				// remove the existing file.
				if err := system.DeleteService(instanceName); err != nil {
					errs = append(errs, err)
					continue
				}
			}

			for _, err := range errs {
				utils.TerminateWithError(err.Error())
			}
		},
	}

	flags = FlagsType{
		force: DeleteCmd.Flags().BoolP("force", "f", false, "force delete packages/instances"),
		name:  DeleteCmd.Flags().StringP("name", "n", "", "name of the deleting instance"),
		all:   DeleteCmd.Flags().BoolP("all", "a", false, "select all packages/instances to delete"),
	}
}

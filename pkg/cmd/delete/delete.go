package delete

import (
	"fmt"
	cmdUtils "go-systemd-docker/pkg/cmd/utils"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"
	"os/exec"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

type Flags struct {
	force *bool
	name  *string
	all   *bool
}

type Delete struct {
	Cmd   *cobra.Command
	Flags Flags
}

func New() *Delete {
	var delete = &Delete{}
	delete.Cmd = &cobra.Command{
		Use:     "delete [flags]",
		Aliases: []string{"rm", "remove"},
		Short:   "Deregister  already register container as Systemd process.",
		Long: `This command deregisters container as Systemd process.
		e.g. sysd delete container-image-name`,
		Args: cobra.MinimumNArgs(0),
		PreRun: func(command *cobra.Command, args []string) {
			if len(args) > 0 && len(*delete.Flags.name) > 0 {
				utils.TerminateWithError("please provide either args[0] or --name not both")
			}

			if !*delete.Flags.all &&
				len(args) == 0 &&
				len(*delete.Flags.name) == 0 {
				utils.TerminateWithError("please provide either args[0] or --name")
			}
		},
		Run: func(command *cobra.Command, args []string) {
			var errs []error
			var instanceNames = []string{*delete.Flags.name}
			if *delete.Flags.all {
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

				if !*delete.Flags.force && svcStatus == service.StatusRunning {
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

				// remove docker-container instance
				drm := exec.Command("docker", "rm", instanceName)
				drmOutput, err := drm.CombinedOutput()
				if err != nil {
					errs = append(errs, fmt.Errorf("docker rm: %s\n%s", drmOutput, err.Error()))
					continue
				}
			}

			if len(errs) > 0 {
				utils.TerminateWithError(fmt.Sprintf("%v", errs))
			}
		},
	}

	delete.Flags = Flags{
		force: delete.Cmd.Flags().BoolP("force", "f", false, "force delete packages/instances"),
		name:  delete.Cmd.Flags().StringP("name", "n", "", "name of the deleting instance"),
		all:   delete.Cmd.Flags().BoolP("all", "a", false, "select all packages/instances to delete"),
	}

	return delete
}

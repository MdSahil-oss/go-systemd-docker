package cmd

import (
	"go-systemd-docker/system"
	"go-systemd-docker/utils"

	"github.com/kardianos/service"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete [flags]",
	Aliases: []string{"rm", "remove"},
	Short:   "Deregister  already register container as Systemd process.",
	Long: `This command deregisters container as Systemd process.
	e.g. sysd delete container-image-name`,
	Args: cobra.RangeArgs(0, 1),
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 && len(*flgs.namePersistentFlag) > 0 {
			utils.Terminate("please provide either args[0] or --name not both")
		}

		if len(args) == 0 && len(*flgs.namePersistentFlag) == 0 {
			utils.Terminate("please provide either args[0] or --name")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var instanceName string = *flgs.namePersistentFlag
		if len(args) > 0 {
			instanceName = args[0]
		}

		svc, err := GetSystemDProcess(instanceName)
		if err != nil {
			utils.Terminate(err.Error())
		}

		svcStatus, err := svc.Status()
		if err != nil {
			utils.Terminate(err.Error())
		}

		if !*flgs.forceFlag && svcStatus == service.StatusRunning {
			utils.Terminate("service is running, please force stop it using '-f'")
		}

		logger, err = svc.Logger(nil)
		if err != nil {
			utils.Terminate(err.Error())
			// log.Fatal(err)
		}

		if svcStatus == service.StatusRunning {
			if err := svc.Stop(); err != nil {
				logger.Error(err)
				utils.Terminate(err.Error())
			}
		}

		if err := svc.Uninstall(); err != nil {
			logger.Error(err)
			utils.Terminate(err.Error())
		}

		// remove the existing file.
		if err := system.DeleteService(instanceName); err != nil {
			utils.Terminate(err.Error())
		}
	},
}

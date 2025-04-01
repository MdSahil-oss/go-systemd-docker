package show

import (
	"fmt"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Flags struct {
	name *string
}

type Show struct {
	Cmd   *cobra.Command
	Flags Flags
}

func New() *Show {
	show := &Show{}

	show.Cmd = &cobra.Command{
		Use:   "show [flags]",
		Short: "Show systemd process configuration for the instance name.",
		// Aliases: []string{"show"},
		Long: `This command List running systemd process (container-image).
		e.g. sysd show sample-instance`,
		Args: cobra.RangeArgs(0, 1),
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) > 0 && len(*show.Flags.name) > 0 {
				utils.TerminateWithError("please provide either args[0] or --name not both")
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			var instanceName string = *show.Flags.name
			if len(args) > 0 {
				instanceName = args[0]
			}

			indexByte, err := os.ReadFile(utils.INDEX_FILE_PATH)
			if err != nil {
				utils.TerminateWithError(err.Error())
				return
			}

			if (len(args) == 0 && len(*show.Flags.name) == 0) || instanceName == "index" {
				fmt.Println(string(indexByte))
				return
			}

			if !system.IsServiceExist(instanceName) {
				utils.TerminateWithError("service doesn't exist")
				return
			}

			index := system.Index{}
			if err = yaml.Unmarshal(indexByte, &index); err != nil {
				utils.TerminateWithError(err.Error())
				return
			}

			for _, svc := range index.Services {
				if svc.Name == instanceName {
					manifestByte, err := os.ReadFile(svc.Path)
					if err != nil {
						utils.TerminateWithError(err.Error())
						return
					}

					fmt.Println(string(manifestByte))
					return
				}
			}

			utils.TerminateWithError("Manifest doesn't exist")
		},
	}

	show.Flags = Flags{
		name: show.Cmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}

	return show
}

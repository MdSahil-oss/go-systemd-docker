package show

import (
	"fmt"
	"go-systemd-docker/pkg/system"
	"go-systemd-docker/pkg/utils"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type FlagsType struct {
	name *string
}

var ShowCmd *cobra.Command
var flags FlagsType

func init() {
	ShowCmd = &cobra.Command{
		Use:     "show [flags]",
		Short:   "Show systemd process configuration for the instance name.",
		Aliases: []string{"ps"},
		Long: `This command List running systemd process (container-image).
		e.g. sysd show sample-instance`,
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

			indexByte, err := os.ReadFile(utils.INDEX_FILE_PATH)
			if err != nil {
				utils.Terminate(err.Error())
				return
			}

			if (len(args) == 0 && len(*flags.name) == 0) || instanceName == "index" {
				fmt.Println(string(indexByte))
				return
			}

			if !system.IsServiceExist(instanceName) {
				utils.Terminate("service doesn't exist")
				return
			}

			index := system.Index{}
			if err = yaml.Unmarshal(indexByte, &index); err != nil {
				utils.Terminate(err.Error())
				return
			}

			for _, svc := range index.Services {
				if svc.Name == instanceName {
					manifestByte, err := os.ReadFile(svc.Path)
					if err != nil {
						utils.Terminate(err.Error())
						return
					}

					fmt.Println(string(manifestByte))
					return
				}
			}

			utils.Terminate("Manifest doesn't exist")
		},
	}

	flags = FlagsType{
		name: ShowCmd.Flags().StringP("name", "n", "", "name of the creating instance"),
	}
}

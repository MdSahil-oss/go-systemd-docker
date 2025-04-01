package create

import (
	"fmt"
	"go-systemd-docker/pkg/utils"
	"strconv"
	"strings"
)

// dockerFlagsCollector serializes all flags with respective values.
func dockerFlagsCollector(flags Flags) []string {
	addtionalCollectiveFlags := []string{}

	if len(*flags.domainName) > 0 {
		addtionalCollectiveFlags = append(addtionalCollectiveFlags, "--domainname")
		addtionalCollectiveFlags = append(addtionalCollectiveFlags, *flags.domainName)
	}

	if len(*flags.entrypoint) > 0 {
		addtionalCollectiveFlags = append(addtionalCollectiveFlags, "--entrypoint")
		addtionalCollectiveFlags = append(addtionalCollectiveFlags, *flags.entrypoint)
	}

	if len(*flags.expose) > 0 {
		for _, val := range *flags.expose {
			// validates if the given value is numeric
			if _, err := strconv.Atoi(val); err != nil {
				utils.TerminateWithError(fmt.Sprintf("port exposing value must be numeric, Found %s", val))
			}

			addtionalCollectiveFlags = append(addtionalCollectiveFlags, "--expose")
			addtionalCollectiveFlags = append(addtionalCollectiveFlags, val)
		}
	}

	if len(*flags.publish) > 0 {
		for _, val := range *flags.publish {
			// validates if the given value is numeric
			if !strings.Contains(val, ":") {
				utils.TerminateWithError(fmt.Sprintf("value must be in key:value format, Found %s", val))
			}

			addtionalCollectiveFlags = append(addtionalCollectiveFlags, "--publish")
			addtionalCollectiveFlags = append(addtionalCollectiveFlags, val)
		}
	}

	if len(*flags.env) > 0 {
		// addtionalCollectiveFlags = strings.Join() // "--expose " + *flags.entrypoint
		for _, val := range *flags.env {
			// validates if the given env value is key=value pair
			if !strings.Contains(val, "=") {
				utils.TerminateWithError(fmt.Sprintf("env variable must be in key=value format, Found %s", val))
			}

			addtionalCollectiveFlags = append(addtionalCollectiveFlags, "--env")
			addtionalCollectiveFlags = append(addtionalCollectiveFlags, val)
		}
	}

	return addtionalCollectiveFlags
}

package utils

import (
	"fmt"
	"go-systemd-docker/pkg/system"
)

func PrintImagesFromIndexService(isvcs []system.IndexService) {
	// Unifying all the images
	set := make(map[string]any)
	for _, element := range isvcs {
		set[element.Image] = struct{}{}
	}

	for k, _ := range set {
		fmt.Println(k)
	}
}

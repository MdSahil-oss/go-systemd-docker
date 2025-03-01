package main

import (
	"fmt"
	"go-systemd-docker/cmd"

	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	v.SetDefault("random", "It should not have been random")
	fmt.Println("before random:", v.GetString("random"))
	cmd.Execute()
}

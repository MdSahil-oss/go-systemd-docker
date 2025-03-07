package main

import (
	"go-systemd-docker/cmd"
)

func main() {
	// v := viper.New()
	// v.SetDefault("random", "It should not have been random")
	// fmt.Println("before random:", v.GetString("random"))
	cmd.Execute()
}

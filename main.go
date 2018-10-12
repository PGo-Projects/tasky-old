package main

import (
	"os"

	"github.com/PGo-Projects/tasky/internal/config"
	"github.com/PGo-Projects/tasky/internal/output"
	"github.com/PGo-Projects/tasky/internal/server"
	"github.com/spf13/cobra"
)

var (
	ServerCmd = &cobra.Command{
		Use: "tasky",
		Run: server.Run,
	}
)

func init() {
	ServerCmd.PersistentFlags().StringVar(&config.Filename, "config", "", "config file (default is config.toml)")
	cobra.OnInitialize(config.Init)
}

func main() {
	if err := ServerCmd.Execute(); err != nil {
		output.Error(err.Error())
		os.Exit(1)
	}
}

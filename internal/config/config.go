package config

import (
	"os"

	"github.com/PGo-Projects/tasky/internal/output"
	"github.com/spf13/viper"
)

var Filename string

func Init() {
	if Filename != "" {
		viper.SetConfigFile(Filename)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}

	if err := viper.ReadInConfig(); err != nil {
		output.Error(err.Error())
		os.Exit(1)
	}
}

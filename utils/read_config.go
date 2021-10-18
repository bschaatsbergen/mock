package utils

import (
	"fmt"
	"os"

	"github.com/bschaatsbergen/mock/model"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ReadConfig reads in the config file and returns it.
func ReadMockConfig(cfgFile string) model.Config {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		wd, err := os.Getwd()
		cobra.CheckErr(err)

		// Search file in working directory with name ".mock.yaml".
		viper.AddConfigPath(wd)
		viper.SetConfigType("yaml")
		viper.SetConfigFile(".mock.yaml")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("📝 Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
		os.Exit(1) //FIXME: This should be handled by wrapping the command with RunE
	}

	var config model.Config

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1) //FIXME: This should be handled by wrapping the command with RunE
	}

	return config
}

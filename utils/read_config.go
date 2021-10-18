package utils

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// ReadConfig reads in the config file.
func ReadConfig(cfgFile string) {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		wd, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in working directory with name ".mock" (without extension).
		viper.AddConfigPath(wd)
		viper.SetConfigType("yaml")
		viper.SetConfigFile(".mock.yaml")
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("📝 Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

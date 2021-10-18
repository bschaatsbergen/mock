/*
Copyright © 2021 Bruno Schaatsbergen <git@bschaatsbergen.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "initializes the mock process by parsing the '.mock.yaml'",
	Long: `📡 initializes the mock process by parsing the '.mock.yaml'. 
The mock config can be passed by using the '-c' flag, otherwise by default from the current working directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		Exec()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "specifies the config file (defaults to '.mock.yaml'")

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
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

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("📝 Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Exec() {
	fmt.Println("Foo")
}

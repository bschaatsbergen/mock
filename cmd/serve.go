package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bschaatsbergen/mock/utils"
)

var cfgFile string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "initializes the mock process by parsing the '.mock.yaml'",
	Long: `📡 initializes the mock process by parsing the '.mock.yaml'. 
The mock config can be passed by using the '-c' flag, otherwise by default from the current working directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "specifies the config file (defaults to '.mock.yaml'")
}

func Serve() {
	utils.ReadMockConfig(cfgFile)
}

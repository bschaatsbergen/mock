package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bschaatsbergen/mock/utils"
)

const defaultPort = "7070"

var cfgFile string
var port string

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
	serveCmd.Flags().StringVarP(&port, "port", "p", defaultPort, "binds a given port to mock")
	serveCmd.Flags().StringVarP(&cfgFile, "config", "c", "", "overwrite the .mock.yaml from the working directory")
}

func Serve() {
	conf := utils.ReadMockConfig(cfgFile)

	if conf.Port != "" {
		port = conf.Port
	}

	utils.StartServer(conf, port)
}

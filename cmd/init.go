package cmd

import (
	"github.com/spf13/cobra"

	"github.com/bschaatsbergen/mock/utils"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "generates a '.mock.yaml' in the current working directory.",
	Long:  `📝 generates a '.mock.yaml' in the current working directory, use the flags to orchestrate the generation.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.GenerateMockConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

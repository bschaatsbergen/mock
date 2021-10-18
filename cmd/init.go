/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

const mockFileName = ".mock.yaml"

var conf = `# Example .mock.yaml config
Endpoints:
  - Resource: /city/1
    Method: GET
    Body: { "Id": 1, "Name": "Albuquerque", "Population": 559,374, "State": "New Mexico" }
    StatusCode: 200

  - Resource: /city
    Method: POST
    Body: { "Name": "Albuquerque", "Population": 559,374, "State": "New Mexico" }
    statusCode: 200

  - Resource: /city/1
    Method: PUT
    Body: { "Population": 601,255 }
    StatusCode: 204

  - Resource: /city/1
    Method: DELETE
    StatusCode: 204
`

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "generates a '.mock.yaml' in the current working directory.",
	Long:  `📝 generates a '.mock.yaml' in the current working directory, use the flags to orchestrate the generation.`,
	Run: func(cmd *cobra.Command, args []string) {
		GenerateMockConfig()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func GenerateMockConfig() {
	dir, _ := os.Getwd()
	if _, err := os.Stat(mockFileName); os.IsNotExist(err) {
		err := ioutil.WriteFile(mockFileName, []byte(conf), 0755)
		if err != nil {
			fmt.Printf("Unable to write file: %v", err)
		} else {
			fmt.Printf("🎉 Generated '.mock.yaml' in: %v\n", dir)
		}
	} else {
		fmt.Printf("'.mock.yaml' file already exists in: %v\n", dir)
	}
}

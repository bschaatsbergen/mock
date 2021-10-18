package utils

import (
	"fmt"
	"io/ioutil"
	"os"
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

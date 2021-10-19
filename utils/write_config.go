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
    Response: '{ "Id": 1, "Name": "Albuquerque", "Population": 559.374, "State": "New Mexico" }'
    StatusCode: 200

  - Resource: /city
    Method: POST
    Response: '{ "Name": "Albuquerque", "Population": 559.374, "State": "New Mexico" }'
    statusCode: 200

  - Resource: /city/1
    Method: PUT
    Response: '{ "Name": "Albuquerque", "Population": 601.255, "State": "New Mexico" }'
    StatusCode: 200

  - Resource: /city/1
    Method: DELETE
    StatusCode: 204
`

func WriteMockConfig() {
	dir, _ := os.Getwd()
	if _, err := os.Stat(mockFileName); os.IsNotExist(err) {
		err := ioutil.WriteFile(mockFileName, []byte(conf), 0755)
		if err != nil {
			fmt.Printf("unable to write file: %v", err)
		} else {
			fmt.Printf("🎉 generated '.mock.yaml' in: %v\n", dir)
		}
	} else {
		fmt.Printf("'.mock.yaml' file already exists in: %v\n", dir)
	}
}

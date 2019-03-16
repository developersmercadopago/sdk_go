package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sdk_go/src/controller"
	"github.com/sdk_go/src/util"
)

func main() {
	// Load the configuration file
	util.Load("config"+string(os.PathSeparator)+"config.json", config)
	fmt.Println(config.AccessToken)
	createPrefences()
}

func createPrefences() {
	//params := "{\n    \"items\": [\n        {\n        \"title\": \"Dummy Item\",\n        \"description\": \"Multicolor Item\",\n        \"quantity\": 1,\n        \"currency_id\": \"ARS\",\n        \"unit_price\": 10.0\n        }\n    ]\n}"
	controller.CreatePrefences()
}

// *****************************************************************************
// Application Settings
// *****************************************************************************

// config the settings variable
var config = &configuration{}

// configuration contains the application settings
type configuration struct {
	AccessToken string `json:"access_token"`
	API         string `json:"api"`
	PublicKey   string `json:"public_key"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

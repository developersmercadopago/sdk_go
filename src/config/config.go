package config

import "encoding/json"

// *****************************************************************************
// Application Settings
// *****************************************************************************

// Config the settings variable
var Config = &configuration{}

// configuration contains the application settings
type configuration struct {
	AccessToken string `json:"access_token"`
	API         string `json:"api"`
	PublicKey   string `json:"public_key"`

	DataBaseUser string `json:"data_base_user"`
	DataBasePass string `json:"data_base_pass"`
	DataBaseName string `json:"data_base_name"`
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

package controller

import (
	"fmt"
	"io/ioutil"

	"github.com/sdk_go/src/api"
	"github.com/sdk_go/src/entity"
)

//CreatePrefences asd
func CreatePrefences() {

	var preference entity.Preference
	preference.LoadPreference()
	res := api.CreatePreferences(&preference)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

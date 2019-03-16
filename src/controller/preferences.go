package controller

import (
	"fmt"
	"io/ioutil"

	"github.com/sdk_go/src/api"
	"github.com/sdk_go/src/model"
)

//CreatePrefences asd
func CreatePrefences() {

	var preference model.Preference
	preference.LoadPreference()

	//fmt.Println(string(preferenceJson))

	/*m := make(map[string]interface{})
	err := json.Unmarshal(preferenceJson, &m)
	if err != nil {
		log.Fatal(err)
	}*/

	res := api.CreatePreferences(&preference)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

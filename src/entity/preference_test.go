package entity

import (
	"testing"
)

func Test_Preference(t *testing.T) {
	var preference Preference
	preference.LoadPreference()

	// res := api.CreatePreferences(&preference)
	// defer res.Body.Close()
	// body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

}

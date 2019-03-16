package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_Preference(t *testing.T) {
	preference := Preference{
		Items: &Items{
			Title:       "Dummy Item",
			Description: "Multicolor Item",
			Quantity:    1,
			CurrencyID:  "ARS",
			UnitPrice:   10.0,
		},
	}

	preferenceJson, _ := json.Marshal(preference)
	fmt.Println(string(preferenceJson))
	
	items := `"items":{"id":"","title":"Dummy Item","description":"Multicolor Item","picture_url":"","category_id":"","quantity":1,"currency_id":"ARS","unit_price":10}`

	if preferenceJson["items"]["title"]


}

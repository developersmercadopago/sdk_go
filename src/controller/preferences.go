package controller

import (
	"fmt"
	"io/ioutil"

	"github.com/sdk_go/src/api"
)

func createPrefences() {
	params := "{\n    \"items\": [\n        {\n        \"title\": \"Dummy Item\",\n        \"description\": \"Multicolor Item\",\n        \"quantity\": 1,\n        \"currency_id\": \"ARS\",\n        \"unit_price\": 10.0\n        }\n    ]\n}"
	res := api.CreatePreferences(params)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}

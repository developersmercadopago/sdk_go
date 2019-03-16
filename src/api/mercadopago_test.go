package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"testing"
)

func Test_CreatePrefences(t *testing.T) {
	params := "{\n    \"items\": [\n        {\n        \"title\": \"Dummy Item\",\n        \"description\": \"Multicolor Item\",\n        \"quantity\": 1,\n        \"currency_id\": \"ARS\",\n        \"unit_price\": 10.0\n        }\n    ]\n}"
	res := CreatePreferences(params)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		log.Fatal(err)
	}

	if m["init_point"] == nil {
		t.Errorf("URL de pagamento não retornada")
	}
}

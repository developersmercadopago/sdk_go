package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sdk_go/src/model"
)

//GetPaymentMethods search the payments methods
func GetPaymentMethods() *http.Response {
	url := "https://api.mercadopago.com/v1/payment_methods?access_token=APP_USR-2090755115109525-061915-434765e937825fb15df81c76d38e13e1-329653108"
	res := ExecGET(url)
	return res
}

// CreatePreferences create the payments preferences
func CreatePreferences(params *model.Preference) *http.Response {

	url := "https://api.mercadopago.com/checkout/preferences?access_token=TEST-4961993466003990-031608-b104e0a70bbfd62cad19894288efefbf-415908242"
	preferenceJSON, _ := json.Marshal(params)

	m := make(map[string]interface{})

	err := json.Unmarshal(preferenceJSON, &m)
	if err != nil {
		log.Fatal(err)
	}

	res := ExecPOST(url, preferenceJSON)
	return res
}

package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetPaymentMethods search the payments methods
func GetPaymentMethods() *http.Response {
	url := "https://api.mercadopago.com/v1/payment_methods?access_token=APP_USR-2090755115109525-061915-434765e937825fb15df81c76d38e13e1-329653108"
	res := Exec(url, "GET", "")
	return res
}

// CreatePreferences create the payments preferences
func CreatePreferences(params string) *http.Response {
	url := "https://api.mercadopago.com/checkout/preferences?access_token=TEST-4961993466003990-031608-b104e0a70bbfd62cad19894288efefbf-415908242"
	res := Exec(url, "POST", params)
	return res
}

// CreatePreferences1 create the payments preferences
func CreatePreferences1(params string) *http.Response {

	url := "https://api.mercadopago.com/checkout/preferences?access_token=TEST-4961993466003990-031608-b104e0a70bbfd62cad19894288efefbf-415908242"

	parametro := "{\n    \"items\": [\n        {\n        \"title\": \"Dummy Item\",\n        \"description\": \"Multicolor Item\",\n        \"quantity\": 1,\n        \"currency_id\": \"ARS\",\n        \"unit_price\": 10.0\n        }\n    ]\n}"

	payload := strings.NewReader(parametro)
	typeRequest := "POST"
	req, _ := http.NewRequest(typeRequest, url, payload)
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	return res
}

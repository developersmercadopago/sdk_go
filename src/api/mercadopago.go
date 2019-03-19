package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sdk_go/src/config"
	"github.com/sdk_go/src/entity"
)

//GetPaymentMethods search the payments methods
func GetPaymentMethods() *http.Response {
	url := fmt.Sprintf("%v/v1/payment_methods?access_token=%v", config.Config.API, config.Config.AccessToken)
	res := ExecGET(url)
	return res
}

// CreatePreferences create the payments preferences
func CreatePreferences(params *entity.Preference) *http.Response {

	url := fmt.Sprintf("%v/checkout/preferences?access_token=%v", config.Config.API, config.Config.AccessToken)
	fmt.Println("url ", url)
	preferenceJSON, _ := json.Marshal(params)

	m := make(map[string]interface{})

	err := json.Unmarshal(preferenceJSON, &m)
	if err != nil {
		log.Fatal(err)
	}

	res := ExecPOST(url, preferenceJSON)
	return res
}

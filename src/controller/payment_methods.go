package controller

import (
	"fmt"
	"io/ioutil"

	"github.com/sdk_go/src/api"
)

func getPaymentsMethods() {

	res := api.GetPaymentMethods()

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}

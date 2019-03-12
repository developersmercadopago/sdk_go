package api

import (
	"net/http"
)

// GetPaymentMethods search the payments methods
func GetPaymentMethods() *http.Response {
	url := "https://api.mercadopago.com/v1/payment_methods?access_token=APP_USR-2090755115109525-061915-434765e937825fb15df81c76d38e13e1-329653108"
	res := ExecGET(url)
	return res
}

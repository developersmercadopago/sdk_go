package api

import (
	"log"
	"net/http"
	"strings"
)

// Exec exec API call
func Exec(url string, typeRequest string, params string) *http.Response {
	payload := strings.NewReader(params)
	log.Printf("Parametros %v e payload %v ", params, payload)
	req, _ := http.NewRequest(typeRequest, url, payload)
	req.Header.Add("content-type", "application/json")
	res, error := http.DefaultClient.Do(req)
	if error != nil {
		panic(error)
	}

	return res
}

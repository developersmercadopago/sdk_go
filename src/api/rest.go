package api

import (
	"bytes"
	"fmt"
	"net/http"
)

// ExecGET exec API call
func ExecGET(url string) *http.Response {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	res, error := http.DefaultClient.Do(req)
	if error != nil {
		panic(error)
	}

	return res
}

// ExecPOST exec API call
func ExecPOST(url string, params []byte) *http.Response {

	fmt.Println(string(params))

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(params))
	req.Header.Add("content-type", "application/json")
	res, error := http.DefaultClient.Do(req)
	if error != nil {
		panic(error)
	}

	return res
}

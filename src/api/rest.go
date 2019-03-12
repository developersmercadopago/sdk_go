package api

import (
	"net/http"
)

// ExecGET exec the GET call
func ExecGET(url string) *http.Response {

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("cache-control", "no-cache")
	res, error := http.DefaultClient.Do(req)
	if error != nil {
		panic(error)
	}

	return res
}

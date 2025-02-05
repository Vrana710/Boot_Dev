package main

import (
	"net/http"
)

func getUserCode(url string) int {
	res, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer res.Body.Close()

	// Return the HTTP status code of the response
	return res.StatusCode
}

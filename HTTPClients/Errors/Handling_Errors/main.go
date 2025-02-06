package main

import (
	"fmt"
	"net/http"
)

func fetchData(url string) (int, error) {
	// Make the HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		// If there's a network error, return 0 and the error message
		return 0, fmt.Errorf("network error: %v", err)
	}
	// Ensure the response body is closed after the function returns
	defer resp.Body.Close()

	// If the status code is not 200, return the status code and an error
	if resp.StatusCode != 200 {
		return resp.StatusCode, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	// If no errors occurred, return the status code and nil for the error
	return resp.StatusCode, nil
}

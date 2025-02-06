package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getResources(path string) []map[string]any {
	fullURL := "https://api.boot.dev" + path // Append path to base URL

	res, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK { // Handle non-200 responses
		fmt.Printf("Error: received status code %d\n", res.StatusCode)
		return nil
	}

	var resources []map[string]any
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&resources)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return nil
	}

	return resources
}
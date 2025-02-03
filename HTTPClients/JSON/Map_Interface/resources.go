package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
)

// getResources sends a GET request to the URL and decodes the JSON response into a slice of maps.
func getResources(url string) ([]map[string]interface{}, error) {
	var resources []map[string]interface{}

	// Send HTTP GET request
	res, err := http.Get(url)
	if err != nil {
		return resources, err
	}
	defer res.Body.Close()

	// Decode the response body into a slice of maps
	err = json.NewDecoder(res.Body).Decode(&resources)
	if err != nil {
		return resources, err
	}

	return resources, nil
}

// logResources iterates over the slice of maps and logs the keys and values, sorted alphabetically by key.
func logResources(resources []map[string]interface{}) {
	var formattedStrings []string

	// Iterate over each map
	for _, resource := range resources {
		// Iterate over key-value pairs in the map
		for key, value := range resource {
			// Format the key-value pair and append it to the slice
			formattedStrings = append(formattedStrings, fmt.Sprintf("Key: %s - Value: %v", key, value))
		}
	}

	// Sort the formatted strings
	sort.Strings(formattedStrings)

	// Print the sorted strings
	for _, str := range formattedStrings {
		fmt.Println(str)
	}
}

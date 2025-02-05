package main

import (
	"fmt"
	"net/http"
)

func deleteUser(baseURL, id, apiKey string) error {
	// Build the full URL for the DELETE request
	fullURL := baseURL + "/" + id

	// Create a new DELETE request
	req, err := http.NewRequest("DELETE", fullURL, nil)
	if err != nil {
		return err
	}

	// Set the required headers
	req.Header.Set("X-API-Key", apiKey)
	req.Header.Set("Content-Type", "application/json")

	// Send the request using http.DefaultClient
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check if the response status code indicates success
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to delete user: %s", res.Status)
	}

	// If successful, return nil (no error)
	return nil
}

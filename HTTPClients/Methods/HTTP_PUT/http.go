package main

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
)

// updateUser updates the user data using the provided baseURL, user ID, API key, and the updated user data.
func updateUser(baseURL, id, apiKey string, data User) (User, error) {
	fullURL := baseURL + "/" + id

	// Marshal user data to JSON
	requestBody, err := json.Marshal(data)
	if err != nil {
		return User{}, err
	}

	// Create a new PUT request
	req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return User{}, err
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	// Check if the status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("failed to update user: %v", resp.Status)
	}

	// Decode the response body into a User struct
	var updatedUser User
	err = json.NewDecoder(resp.Body).Decode(&updatedUser)
	if err != nil {
		return User{}, err
	}

	// Return the updated user
	return updatedUser, nil
}

// getUserById retrieves the user data using the provided baseURL, user ID, and API key.
func getUserById(baseURL, id, apiKey string) (User, error) {
	fullURL := baseURL + "/" + id

	// Create a new GET request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return User{}, err
	}

	// Set the necessary headers
	req.Header.Set("X-API-Key", apiKey)

	// Make the request using the default HTTP client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	// Check if the status code is OK (200)
	if resp.StatusCode != http.StatusOK {
		return User{}, fmt.Errorf("failed to get user: %v", resp.Status)
	}

	// Decode the response body into a User struct
	var user User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return User{}, err
	}

	// Return the user
	return user, nil
}

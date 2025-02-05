package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// createUser makes a POST request to create a new user on the given URL with the provided apiKey
func createUser(url, apiKey string, data User) (User, error) {
	// Step 1: Encode the user data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return User{}, fmt.Errorf("error marshalling data: %v", err)
	}

	// Step 2: Create a new POST request with the JSON data
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return User{}, fmt.Errorf("error creating request: %v", err)
	}

	// Step 3: Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", apiKey)

	// Step 4: Send the request using the default HTTP client
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return User{}, fmt.Errorf("error making the request: %v", err)
	}
	defer res.Body.Close()

	// Step 5: Decode the response body into a User object
	var createdUser User
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&createdUser)
	if err != nil {
		return User{}, fmt.Errorf("error decoding response: %v", err)
	}

	// Return the created user
	return createdUser, nil
}

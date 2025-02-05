package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getUsers(url string) ([]User, error) {
	// Make the GET request to the URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() // Ensure the response body is closed when done

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Decode the JSON into the users slice
	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err
	}

	// Return the decoded users
	return users, nil
}

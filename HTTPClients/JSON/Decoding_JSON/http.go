package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getIssues(url string) ([]Issue, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	defer res.Body.Close()

	var issues []Issue
	if err := json.NewDecoder(res.Body).Decode(&issues); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return issues, nil
}

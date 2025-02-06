package main

import (
	"fmt"
	"net/url"
)

func fetchTasks(baseURL, availability string) []Issue {
	// Determine the limit based on the availability
	var limit int
	switch availability {
	case "Low":
		limit = 1
	case "Medium":
		limit = 3
	case "High":
		limit = 5
	default:
		limit = 0
	}

	// Construct the URL with the sort and limit query parameters
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}
	queryParams := parsedURL.Query()
	queryParams.Add("sort", "estimate") // Sort by estimate
	queryParams.Add("limit", fmt.Sprintf("%d", limit)) // Set the limit based on availability
	parsedURL.RawQuery = queryParams.Encode()

	// Fetch the issues using the modified URL
	return getIssues(parsedURL.String())
}

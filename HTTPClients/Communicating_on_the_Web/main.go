package main

import (
	"fmt"
	"log"
)

func main() {
	issues, err := getIssueData()
	if err != nil {
		log.Fatalf("error getting issue data: %v", err)
	}

	// Don't edit above this line

	// ?
	// Convert byte slice to string
	issuesStr := string(issues)

	// Print the string representation
	fmt.Println(issuesStr)
}

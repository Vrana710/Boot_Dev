package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getIPAddress(domain string) (string, error) {
	url := fmt.Sprintf("https://cloudflare-dns.com/dns-query?name=%s&type=A", domain)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("accept", "application/dns-json")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal the response into DNSResponse struct
	var dnsResponse DNSResponse
	err = json.Unmarshal(body, &dnsResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshaling response: %w", err)
	}

	// Check if there are any answers and return the first IP address
	if len(dnsResponse.Answer) > 0 {
		return dnsResponse.Answer[0].Data, nil
	}

	// If no IP address found, return an empty string and an error
	return "", fmt.Errorf("no IP address found for domain: %s", domain)
}

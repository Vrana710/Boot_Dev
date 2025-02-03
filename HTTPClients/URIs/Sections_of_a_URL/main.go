package main

import (
	"net/url"
)

func newParsedURL(urlString string) ParsedURL {
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return ParsedURL{}
	}

	username := ""
	password := ""
	if parsedUrl.User != nil {
		username = parsedUrl.User.Username()
		password, _ = parsedUrl.User.Password() // Password exists if set
	}

	return ParsedURL{
		protocol: parsedUrl.Scheme,
		username: username,
		password: password,
		hostname: parsedUrl.Hostname(),
		port:     parsedUrl.Port(),
		pathname: parsedUrl.Path,
		search:   parsedUrl.RawQuery,
		hash:     parsedUrl.Fragment,
	}
}

package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const MTBA_API = "https://api-v3.mbta.com"

func newClient() *http.Client {
	// TLS certs should be added here
	return &http.Client{}
}

func newRequest(base, path string, params url.Values) (*http.Request, error) {
	uri, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	uri.Path += path
	params.Add("api_key", os.Getenv("MBTA_API_KEY"))
	uri.RawQuery = params.Encode()

	fmt.Println(uri.String())
	return http.NewRequest("GET", uri.String(), nil)
}

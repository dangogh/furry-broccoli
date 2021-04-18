package main

import (
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

	// if an api key is provided, use it
	// TODO: use something better than env var
	if k := os.Getenv("MBTA_API_KEY"); k != "" {
		params.Add("api_key", k)
	}

	uri.Path += path
	uri.RawQuery = params.Encode()

	return http.NewRequest("GET", uri.String(), nil)
}

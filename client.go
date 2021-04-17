package main

import (
	"net/http"
	"os"
)

type client struct {
	http.Client
	APIKey string
}

func newClient() *client {
	// TODO: get API key from config -- not environment
	return &client{
		APIKey: os.Getenv("MBTA_API_KEY"),
	}
}

func (c *client) Do(req *http.Request) (*http.Response, error) {
	if c.APIKey != "" {
		// Add api key to request if provided
		param := req.URL.Query()
		param.Add("api-key", c.APIKey)
		req.URL.RawQuery = param.Encode()
	}

	return c.Client.Do(req)
}

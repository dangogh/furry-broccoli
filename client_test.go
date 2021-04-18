package main

import (
	"net/url"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const exampleBase = "https://api.example.com"

func Test_newRequest(t *testing.T) {
	apiKey := os.Getenv("MBTA_API_KEY")
	os.Unsetenv("MBTA_API_KEY")
	defer os.Setenv("MBTA_API_KEY", apiKey)

	_, err := newRequest(" \000 bad url ", " bad path ", url.Values{})
	assert.NotNil(t, err)

	req, err := newRequest(exampleBase, "blorfs", url.Values{})
	assert.Nil(t, err)
	assert.NotNil(t, req)

	assert.Equal(t, "/blorfs", req.URL.Path)
	assert.Equal(t, url.Values{}, req.URL.Query())

	// Ensure api_key gets included if provided
	os.Setenv("MBTA_API_KEY", "xxx")
	expVal := url.Values{"api_key": []string{"xxx"}}

	req, err = newRequest(exampleBase, "blorfs", url.Values{})

	assert.Nil(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, "/blorfs", req.URL.Path)
	assert.Equal(t, expVal, req.URL.Query())

	req, err = newRequest(exampleBase, "blorfs",
		url.Values{"fields[stop]": []string{"pluto"}})
	expVal.Add("fields[stop]", "pluto")

	assert.Nil(t, err)
	assert.NotNil(t, req)
	assert.Equal(t, "/blorfs", req.URL.Path)
	assert.Equal(t, expVal, req.URL.Query())
}

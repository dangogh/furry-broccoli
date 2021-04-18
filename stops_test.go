package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var stops_bad = `
{
	"data": [
		{
			"id": "stop1",
			"attributes": {
				"name": "blahblah"
		}
	]
}
`
var stops_test = `
{
	"data": [
		{
			"id": "stop1",
			"attributes": {
				"name": "blahblah"
			}
		},
		{
			"id": "stop2",
			"attributes": {
				"name": "fubar"
			}
		}
	]
}
`

func Test_decodeStops(t *testing.T) {
	stops, err := decodeStops(bytes.NewBuffer([]byte(stops_bad)))

	assert.Nil(t, stops)
	assert.NotNil(t, err)

	stops, err = decodeStops(bytes.NewBuffer([]byte(stops_test)))

	assert.Nil(t, err)
	assert.Len(t, stops, 2)
	assert.Equal(t, "stop1", stops[0].ID)
	assert.Equal(t, "stop2", stops[1].ID)
}

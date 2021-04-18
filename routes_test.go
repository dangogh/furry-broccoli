package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var routes_bad = `
{
	"data": [
		{
			"id": "route1",
			"attributes": {
				"name": "blahblah"
		}
	]
}
`
var routes_test = `
{
	"data": [
		{
			"id": "route1",
			"attributes": {
				"name": "blahblah"
			}
		},
		{
			"id": "route2",
			"attributes": {
				"name": "fubar"
			}
		}
	]
}
`

func Test_decodeRoutes(t *testing.T) {
	routes, err := decodeRoutes(bytes.NewBuffer([]byte(routes_bad)))

	assert.Nil(t, routes)
	assert.NotNil(t, err)

	routes, err = decodeRoutes(bytes.NewBuffer([]byte(routes_test)))

	assert.Nil(t, err)
	assert.Len(t, routes, 2)
	assert.Equal(t, "route1", routes[0].ID)
	assert.Equal(t, "route2", routes[1].ID)
}

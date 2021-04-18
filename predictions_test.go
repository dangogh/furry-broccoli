package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var predictions_bad = `
{
	"data": [
		{
			"id": "prediction1",
			"attributes": {
				"name": "blahblah"
		}
	]
}
`
var predictions_test = `
{
	"data": [
		{
			"id": "prediction1",
			"attributes": {
				"name": "blahblah"
			}
		},
		{
			"id": "prediction2",
			"attributes": {
				"name": "fubar"
			}
		}
	]
}
`

func Test_decodePredictions(t *testing.T) {
	predictions, err := decodePredictions(bytes.NewBuffer([]byte(predictions_bad)))

	assert.Nil(t, predictions)
	assert.NotNil(t, err)

	predictions, err = decodePredictions(bytes.NewBuffer([]byte(predictions_test)))

	assert.Nil(t, err)
	assert.Len(t, predictions, 2)
	assert.Equal(t, "prediction1", predictions[0].ID)
	assert.Equal(t, "prediction2", predictions[1].ID)
}

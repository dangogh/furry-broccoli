package main

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var decode_test = `
{
	"data": [
		{
			"id": "route1",
			"attributes": {
				"long_name": "blahblah"
			}
		}
	]
}
`

func Test_decodeResult(t *testing.T) {
	res, err := decodeResult(bytes.NewBuffer([]byte(decode_test)))

	assert.Nil(t, err)
	assert.Len(t, res.Data, 1)
	assert.Equal(t, "route1", res.Data[0].ID)

	var m map[string]string
	err = json.Unmarshal(res.Data[0].Attributes, &m)
	assert.Nil(t, err)
	assert.Equal(t, map[string]string{"long_name": "blahblah"}, m)
}

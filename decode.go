package main

import (
	"encoding/json"
	"io"
)

type AttrMap struct {
	ID         string          `json:"id"`
	Attributes json.RawMessage `json:"attributes,omitempty"`
}

type Result struct {
	Data []AttrMap `json:"data"`
}

// decodeResult unmarshals the common part of a response. The Attributes portion
// is left as a json string for subsequent unmarshaling by the caller.
func decodeResult(rdr io.Reader) (Result, error) {
	var res Result

	dec := json.NewDecoder(rdr)
	err := dec.Decode(&res)

	return res, err
}

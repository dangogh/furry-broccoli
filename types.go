package main

import "encoding/json"

type AttrMap struct {
	ID         string          `json:"id"`
	Attributes json.RawMessage `json:"attributes,omitempty"`
}

type Result struct {
	Data []AttrMap `json:"data"`
}

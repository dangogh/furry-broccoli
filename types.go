package main

type Attribute struct {
	LongName string `json:"long_name"`
}

type AttrMap struct {
	ID         string    `json:"id"`
	Attributes Attribute `json:"attributes"`
}

type Result struct {
	Data []AttrMap `json:"data"`
}

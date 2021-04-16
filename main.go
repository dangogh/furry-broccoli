package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const MTBA_API = "https://api-v3.mbta.com"

func main() {
	req, err := http.NewRequest("GET", MTBA_API+"/routes", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()

	// TODO: consts for route type to create csv string
	q.Add("filter[type]", "0,1")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(resp.Body)

	var m map[string]interface{}

	err = dec.Decode(&m)
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent(" ", "  ")

	err = enc.Encode(m)
}

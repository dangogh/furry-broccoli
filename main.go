package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

const MTBA_API = "https://api-v3.mbta.com"

func routesRequest(base string) (*http.Request, error) {
	req, err := http.NewRequest("GET", base+"/routes", nil)
	if err != nil {
		return req, err
	}

	q := req.URL.Query()

	// TODO: add consts for vehicle type to for csv string
	// only include light rail (0) and heavy rail (1)
	q.Add("filter[type]", "0,1")

	q.Add("fields[route]", "long_name")
	return req, nil
}

func main() {
	client := http.Client{}

	req, err := routesRequest(MTBA_API)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(req.URL.RawQuery)
	resp, err := client.Do(req)
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

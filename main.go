package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const MTBA_API = "https://api-v3.mbta.com"

func newRequest(base, path string, params url.Values) (*http.Request, error) {
	uri, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	uri.Path += path
	uri.RawQuery = params.Encode()

	return http.NewRequest("GET", uri.String(), nil)
}

func main() {
	c := client{}

	routes, err := c.getRoutes()
	if err != nil {
		log.Fatal(err)
	}

	route := chooseRoute(routes)

	//predictions := getPredictions(route)
	fmt.Println("you chose ", route)

}

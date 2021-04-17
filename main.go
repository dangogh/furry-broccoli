package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

const MTBA_API = "https://api-v3.mbta.com"

func newRequest(base, path string, params url.Values) (*http.Request, error) {
	uri, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	uri.Path += path
	uri.RawQuery = params.Encode()

	fmt.Println(uri.String())
	return http.NewRequest("GET", uri.String(), nil)
}

func main() {
	routes, err := getRoutes()
	if err != nil {
		log.Fatal(err)
	}

	route, err := chooseRoute(routes)
	if err != nil {
		log.Fatal(err)
	}

	stops, err := getStops(route)
	if err != nil {
		log.Fatal(err)
	}

	stop, err := chooseStop(stops)
	if err != nil {
		log.Fatal(err)
	}

	dir, err := chooseDirection(route.DirectionNames)
	if err != nil {
		log.Fatal(err)
	}

	predictions, err := getPredictions(route, stop, dir)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	for _, p := range predictions {
		fmt.Printf("%v, %d\n", p, p.DepartureTime.Sub(now))
	}

	fmt.Println("you chose ", route, route.DirectionNames[dir], stop)

}

package main

import (
	"fmt"
	"log"
	"time"
)

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

	// max diff
	now := time.Now()
	t := time.Time{}

	for _, p := range predictions {
		if t.IsZero() || p.DepartureTime.Sub(now) < t.Sub(now) {
			t = p.DepartureTime
		}
	}

	t = t.In(time.Local)

	fmt.Printf("The next train %s from %s will depart at %s -- %d minutes from now.\n",
		route.DirectionNames[dir], stop.Name, t.Format(time.Kitchen), t.Sub(now)/time.Minute)
}

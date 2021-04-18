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

	dirIdx, err := chooseDirection(route.DirectionNames)
	if err != nil {
		log.Fatal(err)
	}
	dir := route.DirectionNames[dirIdx]

	predictions, err := getPredictions(route, stop, dirIdx)
	if err != nil {
		log.Fatal(err)
	}

	// use the default zero-time as a sentinel that no predictions recorded, yet
	t := time.Time{}

	for _, p := range predictions {
		if p.DepartureTime.Before(t) || t.IsZero() {
			t = p.DepartureTime
		}
	}

	fmt.Println(reportNextDeparture(route, stop, dir, t))
}

func reportNextDeparture(route Route, stop Stop, dir string, t time.Time) string {
	if t.IsZero() {
		return fmt.Sprintf("No predicted departures on route %s %s from %s.",
			route.LongName, dir, stop.Name)
	}

	t = t.In(time.Local)
	mins := time.Until(t) / time.Minute

	return fmt.Sprintf("The next train on %s %s from %s will depart at %s -- %d minutes from now.\n",
		route.LongName, dir, stop.Name, t.Format(time.Kitchen), mins)
}

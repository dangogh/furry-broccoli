package main

import (
	"errors"

	"github.com/manifoldco/promptui"
)

func chooseRoute(routes []Route) (Route, error) {
	if len(routes) == 0 {
		return Route{}, errors.New("no routes found")
	}
	names := make([]string, len(routes))
	for i, r := range routes {
		names[i] = r.LongName
	}

	prompt := promptui.Select{
		Label: "Select a route",
		Items: names,
	}

	var (
		pos int
		err error
	)

	// until we get a valid result
	for {
		pos, _, err = prompt.Run()
		if err == nil {
			break
		}
	}

	return routes[pos], nil
}

func chooseStop(stops []Stop) (Stop, error) {
	if len(stops) == 0 {
		return Stop{}, errors.New("no stops found")
	}
	names := make([]string, len(stops))
	for i, s := range stops {
		names[i] = s.Name
	}

	prompt := promptui.Select{
		Label: "Select a stop",
		Items: names,
	}

	var (
		pos int
		err error
	)

	// until we get a valid result
	for {
		pos, _, err = prompt.Run()
		if err == nil {
			break
		}
	}

	return stops[pos], nil
}

func chooseDirection(dirs []string) (int, error) {
	if len(dirs) == 0 {
		return 0, errors.New("no directions listed")
	}
	prompt := promptui.Select{
		Label: "Select a direction",
		Items: dirs,
	}

	var (
		pos int
		err error
	)

	// until we get a valid result
	for {
		pos, _, err = prompt.Run()
		if err == nil {
			break
		}
	}

	return pos, nil
}

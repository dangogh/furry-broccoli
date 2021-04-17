package main

import (
	"github.com/manifoldco/promptui"
)

func chooseRoute(routes []Route) Route {
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

	return routes[pos]
}

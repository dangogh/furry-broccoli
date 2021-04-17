package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
)

func numberInRange(min, max int) func(string) error {
	return func(in string) error {
		i, err := strconv.Atoi(in)
		if err != nil {
			return errors.New("not a number")
		}
		if i < min || i > max {
			return errors.New("out of range")
		}
		return nil
	}
}

func chooseRoute(routes []Route) Route {
	for i, r := range routes {
		fmt.Printf("%5d) %s\n", i, r.LongName)
	}

	prompt := promptui.Prompt{
		Label:    "Choose route by number",
		Validate: numberInRange(0, len(routes)-1),
	}

	var result string
	var err error

	// until we get a valid result
	for {
		result, err = prompt.Run()
		if err == nil {
			break
		}
	}

	// we know it's an int in range -- no need to check error
	i, _ := strconv.Atoi(result)
	return routes[i]
}

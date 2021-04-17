package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/manifoldco/promptui"
)

const MTBA_API = "https://api-v3.mbta.com"

type Attribute struct {
	Name string `json:"long_name"`
}

type AttrMap struct {
	ID         string    `json:"id"`
	Attributes Attribute `json:"attributes"`
}

type Result struct {
	Data []AttrMap `json:"data"`
}

type client struct {
	http.Client
}

func newRequest(base, path string, params url.Values) (*http.Request, error) {
	uri, err := url.Parse(base)
	if err != nil {
		return nil, err
	}

	uri.Path += path
	uri.RawQuery = params.Encode()

	fmt.Printf("%++v\n", uri)
	return http.NewRequest("GET", uri.String(), nil)
}

func (c *client) getRouteNames() []string {
	// Only need the long_name for now and
	// filter by route type light rail or heavy rail
	params := url.Values{}
	params.Add("fields[route]", "long_name")
	params.Add("filter[type]", "0,1")

	req, err := newRequest(MTBA_API, "routes", params)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(resp.Body)

	var res Result

	err = dec.Decode(&res)
	if err != nil {
		log.Fatal(err)
	}

	routes := make([]string, 0, len(res.Data))

	for _, attribs := range res.Data {
		routes = append(routes, attribs.Attributes.Name)
	}

	return routes
}

func main() {
	client := client{}

	route := client.chooseRoute()

	fmt.Println("you chose ", route)
}

func (c *client) chooseRoute() string {
	routes := c.getRouteNames()

	for i, r := range routes {
		fmt.Printf("%5d) %s\n", i, r)
	}

	prompt := promptui.Prompt{
		Label: "Choose route by number",
		Validate: func(in string) error {
			i, err := strconv.Atoi(in)
			if err != nil {
				return errors.New("not a number")
			}
			if i < 0 || i > len(routes)-1 {
				return errors.New("out of range")
			}
			return nil
		},
	}

	var result string
	var err error
	for {
		result, err = prompt.Run()
		if err == nil {
			break
		}
	}

	i, _ := strconv.Atoi(result)
	return routes[i]
}

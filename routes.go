package main

import (
	"encoding/json"
	"io"
	"net/url"
)

type Route struct {
	ID             string
	LongName       string   `json:"long_name"`
	DirectionNames []string `json:"direction_names"`
}

// getRoutes returns a list of routes with the specified types
// TODO: pass in list of types to include and generate filter
func getRoutes() ([]Route, error) {
	// TODO: add API key in more central place
	params := url.Values{}

	// Only get attrs we need and filter route by type light rail or heavy rail
	params.Add("fields[route]", "long_name,direction_names")
	params.Add("filter[type]", "0,1")

	req, err := newRequest(MTBA_API, "routes", params)
	if err != nil {
		return nil, err
	}

	resp, err := newClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return decodeRoutes(resp.Body)
}

func decodeRoutes(rdr io.Reader) ([]Route, error) {
	res, err := decodeResult(rdr)
	if err != nil {
		return nil, err
	}

	routes := make([]Route, 0, len(res.Data))

	for _, data := range res.Data {
		r := Route{ID: data.ID}
		err := json.Unmarshal(data.Attributes, &r)
		if err != nil {
			return routes, err
		}

		routes = append(routes, r)
	}

	return routes, nil
}

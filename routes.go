package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type client struct {
	http.Client
}

type Route struct {
	ID       string
	LongName string
}

func (c *client) getRoutes() ([]Route, error) {
	// Only need the long_name for now and
	// filter by route type light rail or heavy rail
	params := url.Values{}
	params.Add("fields[route]", "long_name")
	params.Add("filter[type]", "0,1")

	req, err := newRequest(MTBA_API, "routes", params)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)

	var res Result

	err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	routes := make([]Route, 0, len(res.Data))

	for _, attribs := range res.Data {
		routes = append(routes,
			Route{
				LongName: attribs.Attributes.LongName,
				ID:       attribs.ID,
			},
		)
	}

	return routes, nil
}

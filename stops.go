package main

import (
	"encoding/json"
	"net/url"
)

type Stop struct {
	ID   string
	Name string `json:"name"`
}

// getStops returns a list of stops on the specified route
// TODO: pass in list of types to include and generate filter
func getStops(r Route) ([]Stop, error) {
	// TODO: add API key in more central place
	params := url.Values{}

	// Only get attrs we need and filter route by type light rail or heavy rail
	params.Add("fields[stop]", "name")
	params.Add("filter[route]", r.ID)

	req, err := newRequest(MTBA_API, "stops", params)
	if err != nil {
		return nil, err
	}

	resp, err := newClient().Do(req)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)

	var res Result

	err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	stops := make([]Stop, 0, len(res.Data))

	for _, data := range res.Data {
		s := Stop{ID: data.ID}
		err := json.Unmarshal(data.Attributes, &s)
		if err != nil {
			return nil, err
		}

		stops = append(stops, s)
	}

	return stops, nil
}

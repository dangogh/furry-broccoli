package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"time"
)

type Prediction struct {
	ID            string
	DepartureTime time.Time `json:"departure_time"`
}

// getPredictions returns list of predictions for given route
func getPredictions(r Route, s Stop, directionID int) ([]Prediction, error) {
	params := url.Values{}
	params.Add("fields[prediction]", "departure_time")
	params.Add("filter[route]", r.ID)
	params.Add("filter[stop]", s.ID)
	params.Add("filter[direction_id]", fmt.Sprint(directionID))

	req, err := newRequest(MTBA_API, "predictions", params)
	if err != nil {
		return nil, err
	}

	resp, err := newClient().Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return decodePredictions(resp.Body)
}

func decodePredictions(rdr io.Reader) ([]Prediction, error) {
	res, err := decodeResult(rdr)
	if err != nil {
		return nil, err
	}

	predictions := make([]Prediction, 0, len(res.Data))

	for _, data := range res.Data {
		p := Prediction{ID: data.ID}
		err := json.Unmarshal(data.Attributes, &p)
		if err != nil {
			return predictions, err
		}

		predictions = append(predictions, p)
	}

	return predictions, nil
}

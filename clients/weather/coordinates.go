package weather

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (c *Client) Geocoding(cityName string) ([]Coordinates, error) {
	q := url.Values{}
	q.Set("q", cityName)
	q.Set("appid", c.apiKey)

	data, err := c.doRequest(q)
	if err != nil {
		return nil, err
	}
	// log.Print(string(data))
	var coordinates []Coordinates

	if err := json.Unmarshal(data, &coordinates); err != nil {
		return nil, fmt.Errorf("can't decode in struct Coordinates: %w", err)
	}

	return coordinates, nil
}

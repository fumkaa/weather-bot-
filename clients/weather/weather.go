package weather

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

func (c *Client) Weather(lat float64, lon float64) (CurrentWeather, error) {
	q := url.Values{}
	q.Set("lang", "ru")
	q.Set("appid", c.apiKey)
	q.Set("lat", strconv.FormatFloat(lat, 'f', -1, 64))
	q.Set("lon", strconv.FormatFloat(lon, 'f', -1, 64))
	q.Set("units", "metric")

	data, err := c.doRequest(q)
	if err != nil {
		return CurrentWeather{}, err
	}

	var weather CurrentWeather

	if err := json.Unmarshal(data, &weather); err != nil {
		return CurrentWeather{}, fmt.Errorf("can't decode in CurrentWeather: %w", err)
	}
	return weather, nil
}

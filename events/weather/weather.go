package weather

import (
	"errors"
	"fmt"

	"github.com/fumkaa/weather-bot-/clients/weather"
)

type Weather struct {
	geo weather.Client
	wea weather.Client
}

type Meta struct {
	Wea         string
	Description string
	Temp        float64
}

var ErrEmptyResponseGeocoding = errors.New("empty geocoding") //fmt.Errorf("empty geocoding")

func New(geo weather.Client, wea weather.Client) *Weather {
	return &Weather{
		geo: geo,
		wea: wea,
	}

}

func (w *Weather) Weather(cityName string) (Meta, error) {
	geo, err := w.geo.Geocoding(cityName)
	// log.Print(geo)
	if err != nil {
		return Meta{}, fmt.Errorf("can't get coordinate: %w", err)
	}

	if len(geo) == 0 {
		return Meta{}, ErrEmptyResponseGeocoding
	}

	wea, err := w.wea.Weather(geo[0].Latitude, geo[0].Longitude)
	if err != nil {
		return Meta{}, fmt.Errorf("can't get weather: %w", err)
	}

	res := Meta{
		Wea:         wea.Weather[0].Wea,
		Description: wea.Weather[0].Description,
		Temp:        wea.Main.Temp,
	}

	return res, nil
}

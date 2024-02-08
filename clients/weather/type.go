package weather

type CurrentWeather struct {
	Weather []Weather   `json:"weather"`
	Main    Temperature `json:"main"`
}

type Weather struct {
	Wea         string `json:"main"`
	Description string `json:"description"`
}

type Temperature struct {
	Temp float64 `json:"temp"`
}

type Coordinates struct {
	//широта
	Latitude float64 `json:"lat"`
	//долгоа
	Longitude float64 `json:"lon"`
}

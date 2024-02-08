package telegram

import (
	"errors"
	"fmt"
	"log"
	"strings"

	weaClient "github.com/fumkaa/weather-bot-/clients/weather"
	"github.com/fumkaa/weather-bot-/events/weather"
)

// telegram
const (
	startCmd = "/start"
	helpCmd  = "/help"
)

// weather
const (
	hostWeather = "api.openweathermap.org"
	basePathWea = "/data/2.5/weather"
	basePathGeo = "/geo/1.0/direct"
)

func (p *Processor) doCmd(text string, chatId int, userName string, weatherApi string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, userName)
	wea := weather.New(
		*weaClient.New(hostWeather, basePathGeo, weatherApi),
		*weaClient.New(hostWeather, basePathWea, weatherApi),
	)

	data, err := wea.Weather(text)
	// log.Print(data)
	if errors.Is(err, weather.ErrEmptyResponseGeocoding) {
		return p.tg.SendMessage(chatId, msgNotCorrectCityName)
	}

	switch text {
	case helpCmd:
		return p.sendHelp(chatId)
	case startCmd:
		return p.sendHello(chatId)
	default:
		return p.tg.SendMessage(chatId, fmt.Sprintf("Погода: %s\nописание: %s\nтемпература: %g", data.Wea, data.Description, data.Temp))
	}
}

func (p *Processor) sendHello(chatId int) error {
	return p.tg.SendMessage(chatId, msgHello)
}

func (p *Processor) sendHelp(chatId int) error {
	return p.tg.SendMessage(chatId, msgHelp)
}

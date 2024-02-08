package main

import (
	"flag"
	"log"

	tgClient "github.com/fumkaa/weather-bot-/clients/telegram"
	event_consumer "github.com/fumkaa/weather-bot-/consumer/event-consumer"
	"github.com/fumkaa/weather-bot-/events/telegram"
)

const (
	hostTg    = "api.telegram.org"
	batchSize = 100
)

func main() {
	token, weaToken := mustToken()
	eventsProcessor := telegram.New(
		*tgClient.New(hostTg, token),
		weaToken,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service stopped: %w", err)
	}
}

// 6777818389:AAGU0ofjeajnD2iOM3mcOHWa5KdlYhLPoIY
func mustToken() (string, string) {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	weaToken := flag.String(
		"weather-token-api",
		"",
		"token for access to weather api",
	)
	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	if *weaToken == "" {
		log.Fatal("weaToken is not specified")
	}
	return *token, *weaToken
}

// fcf91697a72981518b88edf0f218ec90

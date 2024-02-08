package telegram

import (
	"errors"
	"fmt"

	"github.com/fumkaa/weather-bot-/clients/telegram"
	"github.com/fumkaa/weather-bot-/events"
)

type Processor struct {
	tg         telegram.Client
	weatherApi string
	offset     int
	limit      int
}

type Meta struct {
	ChatId   int
	UserName string
}

var (
	ErrUnknowTypeEvent = errors.New("unknow type event")
	ErrUnknowTypeMeta  = errors.New("unknow type meta")
)

func New(tg telegram.Client, weatherApi string) *Processor {
	return &Processor{
		tg:         tg,
		weatherApi: weatherApi,
	}
}

func (p *Processor) Fetch() ([]events.Event, error) {
	updates, err := p.tg.Updates(p.offset, p.limit)
	if err != nil {
		return nil, fmt.Errorf("can't get updates: %w", err)
	}

	if len(updates) == 0 {
		return nil, nil
	}
	res := make([]events.Event, 0, len(updates))

	for _, u := range updates {
		res = append(res, event(u))
	}

	p.offset = updates[len(updates)-1].UpdateId + 1
	return res, nil
}

func (p *Processor) Process(event events.Event) error {
	switch event.Type {
	case events.Message:
		return p.ProcessMessage(event)
	default:
		return fmt.Errorf("can't process message: %w", ErrUnknowTypeEvent)
	}
}

func (p *Processor) ProcessMessage(event events.Event) error {
	meta, err := meta(event)
	if err != nil {
		return fmt.Errorf("can't process message: %w", err)
	}

	if err := p.doCmd(event.Text, meta.ChatId, meta.UserName, p.weatherApi); err != nil {
		return fmt.Errorf("can't doCmd: %w", err)
	}
	return nil
}

func meta(event events.Event) (Meta, error) {
	meta, ok := event.Meta.(Meta)
	if !ok {
		return Meta{}, fmt.Errorf("can't get meta: %w", ErrUnknowTypeMeta)
	}
	return meta, nil
}

func event(update telegram.Update) events.Event {
	updType := fetchType(update)
	res := events.Event{
		Type: updType,
		Text: fetchText(update),
	}

	if updType == events.Message {
		res.Meta = Meta{
			ChatId:   update.Message.Chat.ChatId,
			UserName: update.Message.From.UserName,
		}
	}
	return res
}

func fetchType(update telegram.Update) events.Type {
	if update.Message == nil {
		return events.UnknowType
	}
	return events.Message
}

func fetchText(update telegram.Update) string {
	if update.Message == nil {
		return ""
	}
	return update.Message.Text
}

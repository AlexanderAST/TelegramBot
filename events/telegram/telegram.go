package telegram

import (
	"TelegramBot/Clients/Telegram"
	"TelegramBot/events"
	e "TelegramBot/lib/error"
	"TelegramBot/storage"
)

type Processor struct {
	tg      *Telegram.Client
	offset  int
	storage storage.Storage
}

func New(client *Telegram.Client, storage storage.Storage) *Processor {
	return &Processor{
		tg:      client,
		storage: storage,
	}

}
func (p *Processor) Fetch(limit int) ([]events.Event, error) {
	update, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}
	res := make([]events.Event, 0, len(update))
	for _, u := range update {
		res = append(res, event(u))
	}
}

func event(upd Telegram.Update) events.Event {
	updType := fetchType(upd)
	res := events.Event{
		Type: updType,
		Text: fetchText(upd),
	}
}

func fetchText(upd Telegram.Update) string {
	if upd.Message == nil {
		return ""
	}
	return upd.Message.Text

}

func fetchType(upd Telegram.Update) events.Type {
	if upd.Message == nil {
		return events.Unknown
	}
	return events.Message
}

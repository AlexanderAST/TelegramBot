package telegram

import "TelegramBot/Clients/Telegram"

type Processor struct {
	tg     *Telegram.Client
	offset int
}

func New(client *Telegram.Client) {

}

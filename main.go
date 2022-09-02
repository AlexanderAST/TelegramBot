package main

import (
	"TelegramBot/Clients/Telegram"
	"flag"
	"log"
)

const tgBotHost = "api.telegram.org"

func main() {

	tgClient := Telegram.New(tgBotHost, mustToken())

	//fetcher=fetcher.New()

	//processor=processor.New()

	//consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("Token is not specified")
	}
	return *token
}

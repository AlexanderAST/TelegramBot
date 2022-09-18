package main

import (
	tgClient "TelegramBot/Clients/Telegram"
	event_consumer "TelegramBot/consumer/event-consumer"
	"TelegramBot/events/telegram"
	"TelegramBot/storage/files"
	"flag"
	"log"
)

const tgBotHost = "api.telegram.org"
const storagePatch = "storage"
const batchSize = 100

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePatch),
	)
	log.Print("service started ")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)
	flag.Parse()
	if *token == "" {
		log.Fatal("Token is not specified")
	}
	return *token
}

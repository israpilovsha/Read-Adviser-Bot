package main

import (
	"flag"
	"log"

	tgClient "github.com/israpilovsha/Read-Adviser-Bot/clients/telegram"
	eventconsumer "github.com/israpilovsha/Read-Adviser-Bot/consumer/event-consumer"
	"github.com/israpilovsha/Read-Adviser-Bot/events/telegram"
	"github.com/israpilovsha/Read-Adviser-Bot/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	// consumer.Start(fetcher, processor)
	consumer := eventconsumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("sevice is stoped", err)
	}
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}

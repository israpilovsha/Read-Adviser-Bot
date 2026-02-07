package main

import (
	"flag"
	"log"
)

func main() {
	// token = flags.Get(token)
	t := MustToken()

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New()

	// processor = processor.New()

	// consumer.Start(fetcher, processor)
}

func MustToken() string {
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

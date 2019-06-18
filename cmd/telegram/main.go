package main

import (
	"github.com/nasermirzaei89/telegram"
	"log"
	"os"
)

func main() {
	bot := telegram.New(os.Getenv("TOKEN"))
	res, err := bot.GetUpdates().Do()
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("%+v", res)
}

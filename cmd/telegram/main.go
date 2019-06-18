package main

import (
	"github.com/nasermirzaei89/telegram"
	"log"
	"os"
)

func main() {
	bot := telegram.New(os.Getenv("TOKEN"))
	res, err := bot.GetUpdates()
	if err != nil {
		log.Fatalln(err)
	}

	if res.OK {
		log.Printf("%+v", res.Result)
	} else {
		log.Printf("%d: %s", res.ErrorCode, *res.Description)
	}
}

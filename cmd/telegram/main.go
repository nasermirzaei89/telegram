package main

import (
	"encoding/json"
	"fmt"
	"github.com/nasermirzaei89/telegram"
	"log"
	"os"
)

func main() {
	bot := telegram.New(os.Getenv("TOKEN"))
	res, err := bot.GetMe()
	if err != nil {
		log.Fatalln(err)
	}

	if res.IsOK() {
		b, err := json.MarshalIndent(res.GetUser(), "", "  ")
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s", string(b))
	} else {
		log.Printf("%d: %s", res.Error().GetErrorCode(), res.Error().GetDescription())
	}
}

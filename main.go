package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tl "gopkg.in/telebot.v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	bot, err := tl.NewBot(tl.Settings{
		Token: os.Getenv("TOKEN"),
		Poller: &tl.LongPoller{
			Timeout: 10 * time.Second,
		},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Starting GO bot")

	bot.Handle("/hello", func(c tl.Context) error {
		return c.Send("Hello!")
	})

	bot.Start()
}

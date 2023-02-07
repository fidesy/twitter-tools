package main

import (
	"github.com/fidesy/twitter-tools/internal/service"
	"github.com/fidesy/twitter-tools/internal/telegrambot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	checkError(err)

	s, err := service.New(&service.Config{
		BearerToken: os.Getenv("BEARER_TOKEN"),
		RawProxy:    os.Getenv("PROXY"),
		DBURL:       os.Getenv("DB_URL"),
	})
	checkError(err)
	defer s.CloseDatabase()

	var actions = make(chan string, 10)

	go s.TrackFollowings(actions)

	bot, err := telegrambot.New(os.Getenv("TG_BOT_TOKEN"), s)
	checkError(err)

	err = bot.Start(actions)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"github.com/fidesy/twitter-tools/internal/service"
	"github.com/fidesy/twitter-tools/internal/telegrambot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var twitters = []string{
	"paradigm",
	"FEhrsam",
	"_charlienoyes",
	"RSSH273",
	"a16z",
	"alive_eth",
	"eddylazzarin",
	"Sequoia",
	"Delphi_Digital",
	"blockchaincap",
	"wbrads",
	"brian_armstrong",
	"zxocw",
	"BinanceLabs",
	"tmlee",
	"tferriss",
}

func main() {
	err := godotenv.Load()
	checkError(err)

	s, err := service.New(&service.Config{
		BearerToken: os.Getenv("BEARER_TOKEN"),
		RawProxy:    os.Getenv("PROXY"),
		DBURL:       os.Getenv("DB_URL"),
	})
	checkError(err)

	var actions = make(chan string, 10)

	go s.MonitorFollowings(actions, twitters)

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

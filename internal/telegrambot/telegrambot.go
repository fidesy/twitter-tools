package telegrambot

import (
	"github.com/fidesy/twitter-tools/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

type TelegramBot struct {
	bot  *tgbotapi.BotAPI
	serv *service.Service
}

func New(token string, serv *service.Service) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		bot:  bot,
		serv: serv,
	}, nil
}

func (tg *TelegramBot) Start(actions <-chan string) error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := tg.bot.GetUpdatesChan(updateConfig)

	var chatIDs = make(map[int64]bool)

	go func() {
		for action := range actions {
			for chatID := range chatIDs {
				tg.sendMessage(chatID, action)
			}
		}
	}()

	for update := range updates {
		update := update
		log.Printf("New message from %s: %s\n", update.Message.From.UserName, update.Message.Text)
		go func() {
			chatID := update.Message.Chat.ID

			messageTextSplit := strings.Split(update.Message.Text, " ")
			switch messageTextSplit[0] {
			case "/start":
				chatIDs[chatID] = true
				tg.sendMessage(chatID, "You are successfully subscribed on notifications!")
			case "/followings":
				tg.sendMessage(chatID, "Nothing to send.")

			default:
				tg.sendMessage(chatID, "unknown command")
			}
		}()
	}
	return nil
}

func (tg *TelegramBot) sendMessage(chatID int64, messageText string) {
	msg := tgbotapi.NewMessage(chatID, messageText)
	msg.ParseMode = "HTML"
	_, err_ := tg.bot.Send(msg)
	if err_ != nil {
		log.Printf("Error: %s", err_.Error())
	}
}

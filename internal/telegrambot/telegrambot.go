package telegrambot

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

const greetingMessage = "Welcome to FundInsightsBot. \nHere are available commands:" +
	"\n\n\t/subscribe - get notifications about follows" +
	"\n\t/unsubscribe - disable notifications about follows" +
	"\n\t/top - get top follows for the last 24hours." +
	"\n\nCurrently bot is tracking 200+ twitter accounts."

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
				tg.sendMessage(chatID, greetingMessage)
			case "/subscribe":
				chatIDs[chatID] = true
				tg.sendMessage(chatID, "You are successfully subscribed on notifications!")
			case "/unsubscribe":
				if _, ok := chatIDs[chatID]; ok {
					delete(chatIDs, chatID)
					tg.sendMessage(chatID, "You are successfully unsubscribed on notifications!")
					return
				}
				tg.sendMessage(chatID, "You are not subscribed on notifications!")
			case "/top":
				result, err := tg.serv.GetTopFollowings(context.Background())
				if err != nil {
					tg.sendMessage(chatID, "error:"+err.Error())
					return
				}

				tg.sendMessage(chatID, result)
			case "/deleteabc":
				if len(messageTextSplit) < 2 {
					tg.sendMessage(chatID, "please provide username to delete")
					return
				}
				err := tg.serv.DeleteUser(context.Background(), messageTextSplit[1])
				if err != nil {
					tg.sendMessage(chatID, "error while deleting user: "+err.Error())
					return
				}

				tg.sendMessage(chatID, "Successfully deleted user.")
			case "/addabc":
				if len(messageTextSplit) < 2 {
					tg.sendMessage(chatID, "please provide username to add")
					return
				}

				err := tg.serv.AddUser(context.Background(), messageTextSplit[1])
				if err != nil {
					tg.sendMessage(chatID, "error while adding user: "+err.Error())
					return
				}

				tg.sendMessage(chatID, "Successfully added user.")

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

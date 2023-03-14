package telegrambot

import (
	"context"
	"github.com/fidesy/twitter-tools/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strings"
)

const greetingMessage = "Welcome to FundInsightsBot. \nHere are available commands:" +
	"\n\n\t/subscribe - get notifications about followings" +
	"\n\t/unsubscribe - disable notifications about followings" +
	"\n\t/top - get top followings for the last 24hours."

//"\n\t/add USERNAME - add user to track" +
//"\n\t/delete USERNAME - delete user from tracking"

type TelegramBotConfig struct {
	Token         string
	AdminUsername string
}

type TelegramBot struct {
	config *TelegramBotConfig
	bot    *tgbotapi.BotAPI
	serv   *service.Service
}

func New(config *TelegramBotConfig, serv *service.Service) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		return nil, err
	}

	config.AdminUsername = strings.ToLower(config.AdminUsername)

	return &TelegramBot{
		config: config,
		bot:    bot,
		serv:   serv,
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
				tg.sendMessage(tgbotapi.NewMessage(chatID, action))
			}
		}
	}()

	for update := range updates {
		update := update
		go func() {
			if update.Message == nil {
				return
			}

			if !update.Message.IsCommand() {
				return
			}

			log.Printf("New message from %s: %s\n", update.Message.From.UserName, update.Message.Text)
			chatID := update.Message.Chat.ID
			msg := tgbotapi.NewMessage(chatID, "")

			args := strings.Split(update.Message.Text, " ")[1:]
			switch update.Message.Command() {
			case "start":
				msg.Text = greetingMessage

			case "subscribe":
				chatIDs[chatID] = true
				msg.Text = "You are successfully subscribed on notifications!"

			case "unsubscribe":
				if _, ok := chatIDs[chatID]; ok {
					delete(chatIDs, chatID)
					msg.Text = "You are successfully unsubscribed on notifications!"
					break
				}
				msg.Text = "You are not subscribed on notifications!"

			case "top":
				result, err := tg.serv.GetTopFollowings(context.Background())
				if err != nil {
					msg.Text = "Error:" + err.Error()
					break
				}

				msg.Text = result

			case "add":
				if !tg.isAdmin(update.Message.From.UserName) {
					msg.Text = "You are not allowed to use this method."
					break
				}

				if len(args) < 1 {
					msg.Text = "Please provide username to add."
					break
				}

				err := tg.serv.AddUser(context.Background(), args[0])
				if err != nil {
					msg.Text = "Error while adding user: " + err.Error()
					break
				}

				msg.Text = "Successfully added user."

			case "delete":
				if !tg.isAdmin(update.Message.From.UserName) {
					msg.Text = "You are not allowed to use this method."
					break
				}

				if len(args) < 1 {
					msg.Text = "Please provide username to delete."
					break
				}

				err := tg.serv.DeleteUser(context.Background(), args[0])
				if err != nil {
					msg.Text = "Error while deleting user: " + err.Error()
					break
				}

				msg.Text = "Successfully deleted user."

			default:
				msg.Text = "Unknown command."
			}

			tg.sendMessage(msg)
		}()
	}
	return nil
}

func (tg *TelegramBot) sendMessage(msg tgbotapi.MessageConfig) {
	msg.ParseMode = "HTML"
	_, err := tg.bot.Send(msg)
	if err != nil {
		log.Println("error while sending message:", err.Error())
	}
}

func (tg *TelegramBot) isAdmin(username string) bool {
	return tg.config.AdminUsername == strings.ToLower(username)
}

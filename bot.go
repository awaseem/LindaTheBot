package main

import (
	"log"

	"github.com/awaseem/LindaTheBot/helpers"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const telegramBotKey = "TELEGRAM_KEY"

func main() {
	bot, err := tgbotapi.NewBotAPI(helpers.GetEnvOrElse(telegramBotKey, ""))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Text == "" {
			continue
		}

		if response := HandleMessage(update.Message.Text, *update.Message.From); response == "" {
			continue
		} else {
			bot.Send(helpers.ResponseMessage(update.Message.Chat.ID, response, update.Message.MessageID))
		}
	}
}

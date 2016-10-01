package main

import (
	"log"

	"github.com/awaseem/LindaTheBot/helpers"
	"github.com/awaseem/LindaTheBot/interactions"
	"github.com/awaseem/LindaTheBot/store"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const telegramBotKey = "TELEGRAM_KEY"

func main() {
	// Create store for holding messages
	storeErr := store.Init()
	if storeErr != nil {
		log.Panic(storeErr)
	}
	// Setup telegram polling for updates
	bot, err := tgbotapi.NewBotAPI(helpers.GetEnvOrElse(telegramBotKey, ""))
	if err != nil {
		log.Panic(err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	// Handle updates for anything that happens within the chat
	for update := range updates {
		// continue on any messages that are pictures
		if update.Message == nil || update.Message.Text == "" {
			continue
		}
		saveErr := store.Save(update.Message.From.FirstName, update.Message.From.LastName, update.Message.From.UserName, update.Message.Text, update.Message.Time())
		// Any errors that occur while saving, just ignore them!
		if saveErr != nil {
			log.Printf("Error: Failed to save message!")
		}
		// Send any interactions that are handled
		if response := interactions.HandleMessage(update.Message.Text, *update.Message.From); response != "" {
			bot.Send(helpers.ResponseMessage(update.Message.Chat.ID, response, update.Message.MessageID))
		}
	}
}

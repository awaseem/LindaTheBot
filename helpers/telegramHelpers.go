package helpers

import (
	"log"

	"github.com/awaseem/LindaTheBot/interactions"
	"github.com/awaseem/LindaTheBot/store"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const telegramBotKey = "TELEGRAM_KEY"

var botAPI *tgbotapi.BotAPI

// StartChannel start the channel
func StartChannel() <-chan tgbotapi.Update {
	// Setup telegram polling for updates
	bot, err := tgbotapi.NewBotAPI(GetEnvOrElse(telegramBotKey, ""))
	if err != nil {
		log.Panic(err)
	} else {
		botAPI = bot
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	return updates
}

// ResponseMessage create telegram response message
func ResponseMessage(chatID int64, message string, messageID int) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyToMessageID = messageID
	return msg
}

// HandleUpdates update channels
func HandleUpdates(updates <-chan tgbotapi.Update) {
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
			botAPI.Send(ResponseMessage(update.Message.Chat.ID, response, update.Message.MessageID))
		}
	}
}

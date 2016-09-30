package helpers

import "github.com/go-telegram-bot-api/telegram-bot-api"

// ResponseMessage create telegram response message
func ResponseMessage(chatID int64, message string, messageID int) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, message)
	msg.ReplyToMessageID = messageID
	return msg
}

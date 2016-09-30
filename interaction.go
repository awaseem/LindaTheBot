package main

import (
	"math/rand"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const hiRequest string = "Hi Linda"

var responses = []string{
	" I'm Linda, don't mind me! I'm just listening in...",
	" I'm Linda, I come in peace!",
	" I'm Linda, herro?",
	" I'm Linda, shhhhhhh!!!",
	" I'm Linda, having a great time listening to you guys!",
	" I'm Linda, this isnt creepy right!",
}

func responseWithFirstName(firstName string, message string) string {
	return "Hello " + firstName + message
}

// HandleMessage return a message based on certian requests!
func HandleMessage(message string, user tgbotapi.User) string {
	var response string
	switch message {
	case hiRequest:
		response = responseWithFirstName(user.FirstName, responses[rand.Intn(len(responses))])
	}
	return response
}

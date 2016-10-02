package main

import (
	"log"

	"github.com/awaseem/LindaTheBot/store"
	"github.com/awaseem/LindaTheBot/telegramHelpers"
)

func main() {
	// Create store for holding messages
	storeErr := store.Init()
	if storeErr != nil {
		log.Panic(storeErr)
	}
	updates := telegramHelpers.StartChannel()
	// Handle updates for anything that happens within the chat
	telegramHelpers.HandleUpdates(updates)
}

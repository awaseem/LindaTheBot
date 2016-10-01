package main

import (
	"log"

	"github.com/awaseem/LindaTheBot/helpers"
	"github.com/awaseem/LindaTheBot/store"
)

func main() {
	// Create store for holding messages
	storeErr := store.Init()
	if storeErr != nil {
		log.Panic(storeErr)
	}
	updates := helpers.StartChannel()
	// Handle updates for anything that happens within the chat
	helpers.HandleUpdates(updates)
}

package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var ctx = Ctx

func TelegramHandler(w http.ResponseWriter, r *http.Request) {
	// Clear update before decode
	update := telegram.Update{}
	text := ""
	chatID := 0
	messageID := 0

	// Decode Request in global var update
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		log.Panic(err)
		return
	}

	if update.CallbackQuery != nil {
		text = update.CallbackQuery.Data
		chatID = int(update.CallbackQuery.Message.Chat.ID)
		messageID = update.CallbackQuery.Message.MessageID
	} else {
		text = update.Message.Text
		chatID = int(update.Message.Chat.ID)
		messageID = update.Message.MessageID
	}

	//handle reponse ...
	log.Print(text)
	log.Print(chatID)
	log.Print(messageID)
}

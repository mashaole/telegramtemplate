package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mashaole/telegramtemplate/utils"
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
	go deleteChatHistory(int64(chatID), messageID)

	switch text {
	case "simple":
		params := utils.SendMessageParams{
			ChatID:   int64(chatID),
			Text:     "simple reply",
			Keyboard: StartMenu}
		err := SendSimpleMessage(params)
		if err != nil {
			log.Print(err)
		}
	case "inline":
		params := utils.SendMessageParams{
			ChatID:   int64(chatID),
			Text:     "simple reply",
			Keyboard: StartMenu}
		err := SendInlineMessage(params)
		if err != nil {
			log.Print(err)
		}
	default:
		msg := telegram.NewMessage(int64(chatID), "Please type simple or inline or edit")
		_, err := bot.Send(msg)
		if err != nil {
			log.Print(err)
		}
	}
}

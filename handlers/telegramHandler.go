package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var bot *telegram.BotAPI

func InitTelegram() error {

	if ConfigFile == "" {
		if runtime.GOOS == "windows" {
			ConfigFile = "../config.json"
		} else {
			ConfigFile = "./config.json"
		}
	}

	config, err := GetConfig()
	if err != nil {
		return err
	}

	bot, err = telegram.NewBotAPI(config.BotToken)
	if err != nil {
		return nil
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := telegram.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}

	for update := range updates {
		if update.Message == nil && update.CallbackQuery == nil { // Ignore any non-Message Updates
			continue
		}

		jsonData, err := json.Marshal(update)
		if err != nil {
			return err
		}

		request := &http.Request{
			Body: ioutil.NopCloser(bytes.NewReader(jsonData)),
		}

		var respWriter http.ResponseWriter
		TelegramHandler(respWriter, request)
	}
	return nil
}

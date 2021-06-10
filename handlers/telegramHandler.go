package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api"
)

var bot *telegram.BotAPI

func InitTelegram() error {
	var err error
	bot, err = telegram.NewBotAPI(Token)
	if err != nil {
		return nil
	}

	port := os.Getenv("PORT")

	//In local environment
	if port == "" {
		_, err := bot.RemoveWebhook() // Removes webhook if already set
		if err != nil {
			return err
		}

		bot.Debug = false
		log.Printf("Authorized on account %s", bot.Self.UserName)

		u := telegram.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)
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
	} else {
		//In App Engine
		webURL, err := url.Parse(WebURL)
		if err != nil {
			return err
		}
		_, err = bot.SetWebhook(telegram.WebhookConfig{URL: webURL})
		if err != nil {
			return err
		}
		updates := bot.ListenForWebhook(Endpoint)

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
	}
	return nil
}

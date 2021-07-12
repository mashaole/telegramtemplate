package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/mashaole/telegramtemplate/utils"
)

// Load heavy queries on startup - All transactions take very long to load
var Ctx = context.Background()

var ConfigFile = ""

//UI buttons

var menu = []string{"Menu", "menu"}
var StartMenu = [][]string{
	{"button1", "1"}, {"button2", "2"},
	{"button3", "3"}, menu,
}

//UI buttons

//CreateInlineKeyBoard Creates Inline keyboard after being passed a button texts and their callbacks
func CreateInlineKeyBoard(buttons [][]string) telegram.InlineKeyboardMarkup {
	if len(buttons) > 0 {
		var keyboardinline [][]telegram.InlineKeyboardButton

		//This part is an algorithim to add 2 buttons per row else one button will take up with of both columns
		// for to change this change `i+=2` to number of rows needed and repeat should always be 1 less than `i+=2`
		for i := 0; i <= len(buttons)-1; i += 2 {
			var keyboardButtons []telegram.InlineKeyboardButton
			var repeat = 1
			for j := 0; j <= repeat; j++ {
				if i+j == len(buttons) {
					break
				}
				keyboardButtons = append(keyboardButtons, telegram.NewInlineKeyboardButtonData(buttons[i+j][0], buttons[i+j][1]))
			}
			keyboardinline = append(keyboardinline, keyboardButtons)
		}
		return telegram.InlineKeyboardMarkup{
			InlineKeyboard: keyboardinline,
		}
	}
	return telegram.InlineKeyboardMarkup{}
}

//CreateSimpleKeyBoard Creates Simple keyboard after being passed a button texts
func CreateSimpleKeyBoard(buttons [][]string) interface{} {
	if len(buttons) > 0 {
		var keyboard [][]telegram.KeyboardButton
		for _, texts := range buttons {
			var keyboardButtons []telegram.KeyboardButton

			for _, text := range texts {
				z := telegram.NewKeyboardButton(text)
				keyboardButtons = append(keyboardButtons, z)
			}
			keyboard = append(keyboard, keyboardButtons)
		}
		return telegram.ReplyKeyboardMarkup{
			Keyboard:       keyboard,
			ResizeKeyboard: true}
	}
	return nil
}

//SendMessage method for sending inline reply
func SendSimpleMessage(params utils.SendMessageParams) error {
	msg := telegram.NewMessage(params.ChatID, params.Text)
	msg.ReplyMarkup = CreateSimpleKeyBoard(params.Keyboard)
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

//SendMessage method for sending inline reply
func SendInlineMessage(params utils.SendMessageParams) error {
	msg := telegram.NewMessage(params.ChatID, params.Text)
	msg.ReplyMarkup = CreateInlineKeyBoard(params.Keyboard)
	_, err := bot.Send(msg)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

//EditMessage method for editing inline message
func EditMessage(params utils.EditMessageParams) error {
	keyboard := CreateInlineKeyBoard(params.Keyboard)
	msg := telegram.NewEditMessageText(params.ChatID, params.MessageID, params.Text)
	msg.ParseMode = params.ParseMode
	msg.ReplyMarkup = &keyboard
	_, err := bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

//deleteChatHistory method for deleting Previous messages in chat
func deleteChatHistory(chatID int64, messageID int) {
	for i := 1; i < 10; i++ {
		clearChat := telegram.NewDeleteMessage(chatID, messageID-i)
		_, err := bot.Send(clearChat)
		if err != nil {
			log.Print(err)
		}
	}
}

//GetConfig gets bot token from config.json
func GetConfig() (*utils.Config, error) {
	jsonFile, err := os.Open(ConfigFile)
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err := jsonFile.Close()
		if err != nil {
			log.Print(err)
		}
	}()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var file *utils.Config
	err = json.Unmarshal([]byte(byteValue), &file)
	if err != nil {
		fmt.Println(err)
	}
	if file.BotToken == "" {
		return nil, fmt.Errorf("telegram bot token is not declared in the file " + ConfigFile)
	}
	return file, nil
}

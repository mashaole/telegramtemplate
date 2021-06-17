package main

import (
	"fmt"

	handlers "github.com/mashaole/telegramtemplate/handlers"
)

func main() {
	fmt.Println("Telegram bot started ...")
	err := handlers.InitTelegram()
	if err != nil {
		fmt.Printf("could not initialize telegram bot: %v", err)
		return
	}
}

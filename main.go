package main

import (
	"fmt"

	handlers "github.com/Celbux/celbuxStats-telegram-bot/handlers"
)

func main() {
	fmt.Println("Telegram bot started ...")
	err := handlers.InitTelegram() //uncomment this line and use command `go run main.go` to run on local workstation`
	if err != nil {
		fmt.Printf("could not initialize telegram bot: %v", err)
		return
	}
}

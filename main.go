package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handlers "github.com/mashaole/telegramtemplate/handlers"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/bsjhdbsakjbdjks", handlers.TelegramHandler)
	go listenAndServe(os.Getenv("PORT"))
	fmt.Println("Telegram bot started ...")
	err := handlers.InitTelegram()
	if err != nil {
		fmt.Printf("could not initialize telegram bot: %v", err)
		return
	}
}

func listenAndServe(port string) {
	// Ensure port
	log.Print("----env port: " + port)
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

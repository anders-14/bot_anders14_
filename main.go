package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	botName := os.Getenv("BOT_NAME")
	botOauth := os.Getenv("BOT_OAUTH")
	channelName := "#" + os.Getenv("CHANNEL")
	fmt.Println(botName, botOauth, channelName)

	client := NewClient(botName, botOauth, channelName)

	client.Connect()
	defer client.Close()
	client.Login()

	client.HandleChat()
}

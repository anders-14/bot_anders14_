package main

import (
	"flag"
	"log"
	"os"

	"github.com/anders-14/bot_anders14_/pkg/client"
	"github.com/joho/godotenv"
)

var (
	channelName   = flag.String("channel", "anders14_", "The main channel to connect to")
	commandPrefix = flag.String("prefix", "!", "What prefix to put in front of commands")
)

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	botName := os.Getenv("BOT_NAME")
	botOauth := os.Getenv("BOT_OAUTH")

	c := client.NewClient(botName, botOauth, *channelName, *commandPrefix)
	defer c.Close()

	c.HandleChat()
}

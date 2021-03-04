package main

import (
	"flag"
	"log"
	"os"
	"strings"

	"github.com/anders-14/bot_anders14_/pkg/client"
	"github.com/joho/godotenv"
)

var (
	// Can either enter a single channel or a commaseperated list of channels (no spaces)
	channelFlag = flag.String("channel", "anders14_", "Channels to connect to, separated by comma")
	prefixFlag  = flag.String("prefix", "!", "Prefix to put in front of commands")
)

func main() {
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	botName := os.Getenv("BOT_NAME")
	botOauth := os.Getenv("BOT_OAUTH")

	channels := strings.Split(*channelFlag, ",")

	c := client.NewClient(botName, botOauth, channels, *prefixFlag)
	defer c.Close()

	c.HandleChat()
}

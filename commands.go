package main

import (
	"fmt"

	"github.com/anders-14/bot_anders14_/pkg/client"
	"github.com/anders-14/bot_anders14_/pkg/joke"
	"github.com/anders-14/bot_anders14_/pkg/rps"
	"github.com/anders-14/bot_anders14_/pkg/trivia"
)

// Commands, map from command name to function
var Commands = map[string]func(c *client.Client, pc *ParsedCommand){

	"color": func(c *client.Client, pc *ParsedCommand) {
		msg := fmt.Sprintf("The color of @%s is %s", pc.user.Name, pc.user.Color)
		c.SendMessage(msg)
	},

	"commands": func(c *client.Client, pc *ParsedCommand) {
		c.SendMessage("Go to bot-anders14-commands.vercel.app for a list of all the commands and what they do")
	},

	"joke": func(c *client.Client, pc *ParsedCommand) {
		joke := joke.FetchJoke().Joke
		c.SendMessage(joke)
	},

	"ping": func(c *client.Client, pc *ParsedCommand) {
		msg := fmt.Sprintf("Pong, @%s", pc.user.Name)
		c.SendMessage(msg)
	},

	"rps": func(c *client.Client, pc *ParsedCommand) {
		if len(pc.args) < 1 {
			return
		}
		usermove := pc.args[0]
		msg := rps.Play(usermove, pc.user.Name)
		c.SendMessage(msg)
	},

	"today": func(c *client.Client, pc *ParsedCommand) {
		trivia := trivia.FetchToday()
		c.SendMessage(trivia)
	},
}

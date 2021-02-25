package main

import (
	"fmt"

	"github.com/anders-14/bot_anders14_/pkg/joke"
	"github.com/anders-14/bot_anders14_/pkg/rps"
	"github.com/anders-14/bot_anders14_/pkg/trivia"
)

// Commands, map from command name to function
var Commands = map[string]func(c *Client, pc *ParsedCommand){

	"color": func(c *Client, pc *ParsedCommand) {
		msg := fmt.Sprintf("The color of @%s is %s", pc.user.displayname, pc.user.color)
		c.SendMessage(msg)
	},

	"commands": func(c *Client, pc *ParsedCommand) {
		c.SendMessage("Go to bot-anders14-commands.vercel.app for a list of all the commands and what they do")
	},

	"joke": func(c *Client, pc *ParsedCommand) {
		joke := joke.FetchJoke().Joke
		c.SendMessage(joke)
	},

	"ping": func(c *Client, pc *ParsedCommand) {
		msg := fmt.Sprintf("Pong, @%s", pc.user.displayname)
		c.SendMessage(msg)
	},

	"rps": func(c *Client, pc *ParsedCommand) {
    if len(pc.args) < 1 {
      return
    }
		usermove := pc.args[0]
    msg := rps.Play(usermove, pc.user.displayname)
    c.SendMessage(msg)
	},

	"today": func(c *Client, pc *ParsedCommand) {
		trivia := trivia.FetchToday()
		c.SendMessage(trivia)
	},
}

package main

import (
	"fmt"
)

var Commands = map[string]func(c *Client, pc *ParsedCommand){
	"color": func(c *Client, pc *ParsedCommand) {
		msg := fmt.Sprintf("The color of @%s is %s", pc.user.displayname, pc.user.color)
		c.SendMessage(msg)
	},
	"commands": func(c *Client, pc *ParsedCommand) {
		c.SendMessage("Go to bot-anders14-commands.vercel.app for a list of all the commands and what they do")
	},
	"joke": func(c *Client, pc *ParsedCommand) {
		joke := FetchJoke().Joke
		c.SendMessage(joke)
	},
	"ping": func(c *Client, pc *ParsedCommand) {
		msg := fmt.Sprintf("Pong, @%s", pc.user.displayname)
		c.SendMessage(msg)
	},
	"rps": func(c *Client, pc *ParsedCommand) {
		userMove := pc.args[0]
		outcome := PlayRPS(userMove)
		var msg string
		if outcome != "" {
			msg = fmt.Sprintf("%s @%s", outcome, pc.user.displayname)
		} else {
			msg = fmt.Sprintf("Invalid move @%s", pc.user.displayname)
		}
		c.SendMessage(msg)
	},
	"today": func(c *Client, pc *ParsedCommand) {
		trivia := FetchTodayTrivia().trivia
		c.SendMessage(trivia)
	},
}

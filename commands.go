package main

import (
	"fmt"
)

type Command struct {
	name string
	exec func(c *Client, pc *ParsedCommand)
}

var Commands = [...]Command{
	{
		name: "commands",
		exec: func(c *Client, pc *ParsedCommand) {
			c.SendMessage("Go to bot-anders14-commands.vercel.app for a list of all the commands and what they do")
		},
	},
	{
		name: "ping",
		exec: func(c *Client, pc *ParsedCommand) {
			msg := fmt.Sprintf("Pong, @%s", pc.user.displayname)
			c.SendMessage(msg)
		},
	},
	{
		name: "color",
		exec: func(c *Client, pc *ParsedCommand) {
			msg := fmt.Sprintf("The color of @%s is %s", pc.user.displayname, pc.user.color)
			c.SendMessage(msg)
		},
	},
	{
		name: "joke",
		exec: func(c *Client, pc *ParsedCommand) {
			joke := FetchJoke().Joke
			c.SendMessage(joke)
		},
	},
	{
		name: "today",
		exec: func(c *Client, pc *ParsedCommand) {
			trivia := FetchTodayTrivia().trivia
			c.SendMessage(trivia)
		},
	},
	{
		name: "rps",
		exec: func(c *Client, pc *ParsedCommand) {
			userMove := pc.args[0]
			outcome := PlayRPS(userMove)
			msg := ""
			if outcome != "" {
				msg = fmt.Sprintf("%s @%s", outcome, pc.user.displayname)
			} else {
				msg = fmt.Sprintf("Not a valid move @%s", pc.user.displayname)
			}
			c.SendMessage(msg)
		},
	},
}

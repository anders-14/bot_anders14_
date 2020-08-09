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
}

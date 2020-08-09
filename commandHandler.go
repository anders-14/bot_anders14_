package main

import (
	"strings"
)

type ParsedCommand struct {
	name string
	args []string
	user User
}

func ParseMessageToCommand(m *Message) *ParsedCommand {
	splitMessage := strings.Split(m.content, " ")
	name := strings.ToLower(splitMessage[0][1:])
	args := splitMessage[1:]
	return &ParsedCommand{
		name: name,
		args: args,
		user: m.user,
	}
}

func HandleCommand(c *Client, pc *ParsedCommand) {
	for _, v := range Commands {
		if v.name == pc.name {
			v.exec(c, pc)
			return
		}
	}
}

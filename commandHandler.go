package main

import (
	"strings"
)

/*
ParsedCommand -> object holding info about the parsed command from the user
*/
type ParsedCommand struct {
	name string
	args []string
	user User
}

/*
ParseMessageToCommand -> parses a Message to a ParsedCommand
*/
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

/*
HandleCommand -> handles the execution of an incomming command
*/
func HandleCommand(c *Client, pc *ParsedCommand) {
	for _, v := range Commands {
		if v.name == pc.name {
			v.exec(c, pc)
			return
		}
	}
}

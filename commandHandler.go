package main

import (
	"fmt"
	"strings"
)

// ParsedCommand -> object holding info about the parsed command from the user
type ParsedCommand struct {
	name string
	args []string
	user User
}

// ParseMessageToCommand -> parses a Message to a ParsedCommand
func ParseMessageToCommand(m *Message) *ParsedCommand {
	prefixLen := len(*commandPrefix)
	splitMessage := strings.Split(m.content, " ")
	name := strings.ToLower(splitMessage[0][prefixLen:])
	args := splitMessage[1:]
	return &ParsedCommand{
		name: name,
		args: args,
		user: m.user,
	}
}

// HandleCommand -> handles the execution of an incomming command
func HandleCommand(c *Client, pc *ParsedCommand) {
	if f, ok := Commands[pc.name]; ok {
		f(c, pc)
	} else {
		fmt.Printf("Error: %s is not a command\n", pc.name)
	}
}

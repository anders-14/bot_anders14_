package main

import (
	"fmt"
	"strings"

	"github.com/anders-14/bot_anders14_/pkg/client"
	"github.com/anders-14/bot_anders14_/pkg/parser"
)

// ParsedCommand -> object holding info about the parsed command from the user
type ParsedCommand struct {
	name string
	args []string
	user parser.User
  Channel string
}

// ParseMessageToCommand -> parses a Message to a ParsedCommand
func ParseMessageToCommand(m *parser.Message) *ParsedCommand {
	prefixLen := len(*commandPrefix)
	splitMessage := strings.Split(m.Content, " ")
	name := strings.ToLower(splitMessage[0][prefixLen:])
	args := splitMessage[1:]
	return &ParsedCommand{
		name: name,
		args: args,
		user: m.User,
    Channel: m.Channel,
	}
}

// HandleCommand -> handles the execution of an incomming command
func HandleCommand(c *client.Client, pc *ParsedCommand) {
	if f, ok := Commands[pc.name]; ok {
		f(c, pc)
	} else {
		fmt.Printf("Error: %s is not a command\n", pc.name)
	}
}

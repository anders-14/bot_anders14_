package main

import (
	"fmt"
	"math/rand"
	"strings"
)

/*
Command -> object holding info about a command
*/
type Command struct {
	name string
	args []string
	user User
}

/*
ParseCommand -> parsing the incomming command
*/
func ParseCommand(message *Message) *Command {
	wordList := strings.Fields(message.content)

	name := strings.ToLower(string(wordList[0][1:]))

	args := wordList[1:]

	return &Command{
		name: name,
		args: args,
		user: message.user}
}

/*
Exec -> executes the command
*/
func (c *Command) Exec(client *Client) {
	if c.name == "commands" {
		client.SendMessage("You can find all commands and a description of them at bot-anders14-commands.vercel.app")
	}

	if c.name == "ping" {
		client.SendMessage("Pong @" + c.user.displayname)
		return
	}

	if c.name == "joke" {
		dadJoke := FetchJoke().Joke
		client.SendMessage(dadJoke)
		return
	}

	if c.name == "rps" {
		if len(c.args) > 0 {
			botMove := PlayRPS(c.args[0])
			client.SendMessage(botMove + " @" + c.user.displayname)
		}
		return
	}

	if c.name == "bot_anders14_" {
		desc := "@bot_anders14_ is a chat bot made by @anders14_. It is written in the programming language go"
		client.SendMessage(desc)
		return
	}

	if c.name == "today" {
		todaysTrivia := FetchTodayTrivia()
		client.SendMessage(todaysTrivia.trivia)
		return
	}

	if c.name == "dice" {
		roll := rand.Intn(6) + 1
		rolledmsg := fmt.Sprintf("You rolled a %d @%s", roll, c.user.displayname)
		client.SendMessage(rolledmsg)
	}
}

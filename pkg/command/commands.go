package command

import (
	"fmt"

	"github.com/anders-14/bot_anders14_/pkg/joke"
	"github.com/anders-14/bot_anders14_/pkg/message"
	"github.com/anders-14/bot_anders14_/pkg/rps"
	"github.com/anders-14/bot_anders14_/pkg/trivia"
)

// Commands, map from command name to function
var Commands = map[string]func(pc *message.Command) string{

	"color": func(pc *message.Command) string {
		return fmt.Sprintf("The color of @%s is %s", pc.User.Name, pc.User.Color)
	},

	"commands": func(pc *message.Command) string {
		return fmt.Sprintf("@%s, go to bot-anders14-commands.vercel.app to find all the commands", pc.User.Name)
	},

	"joke": func(pc *message.Command) string {
		return joke.FetchJoke().Joke
	},

	"ping": func(pc *message.Command) string {
		return fmt.Sprintf("Pong, @%s", pc.User.Name)
	},

	"rps": func(pc *message.Command) string {
		if len(pc.Args) < 1 {
			return ""
		}
		usermove := pc.Args[0]
		return rps.Play(usermove, pc.User.Name)
	},

	"today": func(pc *message.Command) string {
		return trivia.FetchToday()
	},
}

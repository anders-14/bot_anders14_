package command

import (
	"fmt"

	"github.com/anders-14/bot_anders14_/pkg/message"
)

// HandleCommand take a pointer to a message.Command executes the
// command and returns the bots reply
func HandleCommand(pc *message.Command) string {
	if f, ok := Commands[pc.Name]; ok {
		msg := f(pc)
		return msg
	}
	fmt.Printf("Error: %s is not a command\n", pc.Name)
	return ""
}

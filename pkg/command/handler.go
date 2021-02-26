package command

import (
	"fmt"

	"github.com/anders-14/bot_anders14_/pkg/message"
)

func HandleCommand(pc *message.Command) string {
	if f, ok := Commands[pc.Name]; ok {
		msg := f(pc)
    return msg
	} else {
		fmt.Printf("Error: %s is not a command\n", pc.Name)
	}

  return ""
}

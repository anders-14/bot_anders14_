package irc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/textproto"
)

// Read irc and put it onto a channel
func Read(reader io.Reader) <-chan string {
	messages := make(chan string, 100)
	proto := textproto.NewReader(bufio.NewReader(reader))

	go func() {
		defer close(messages)
		for {
			line, err := proto.ReadLine()
			if err != nil {
				log.Fatalf("err: %s", err)
			}

			messages <- line
		}
	}()

	return messages
}

// Send message to irc
func Send(writer io.Writer, msg, channel string) error {
	if msg == "" {
		return fmt.Errorf("Cannot send empty string as message")
	}

	fmt.Fprintf(writer, "PRIVMSG #%s :%s\n", channel, msg)

	return nil
}

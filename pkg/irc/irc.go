package irc

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/textproto"
)

// Read irc and put it onto a channel
func Read(reader io.Reader, raw chan string) {
	proto := textproto.NewReader(bufio.NewReader(reader))

	for {
		line, err := proto.ReadLine()
		if err != nil {
			log.Fatalf("err: %s", err)
		}
		raw <- line
	}
}

// Send message to irc
func Send(writer io.Writer, msg, channel string) error {
	if msg == "" {
		return fmt.Errorf("Cannot send empty string as message")
	}

	fmt.Fprintf(writer, "PRIVMSG #%s :%s\n", channel, msg)

	return nil
}

// Pong the server to keep the connection open
func Pong(writer io.Writer) {
	fmt.Fprintf(writer, "PONG :tmi.twitch.tv\n")
}

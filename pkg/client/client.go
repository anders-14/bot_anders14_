package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"
	"time"

	"github.com/anders-14/bot_anders14_/pkg/command"
	"github.com/anders-14/bot_anders14_/pkg/message"
	"github.com/anders-14/bot_anders14_/pkg/parser"
)

var (
	server string = "irc.chat.twitch.tv"
	port   string = "6667"
)

// Client object holding info about the connection
type Client struct {
	Nick          string
	Channel       string
	CommandPrefix string
	Conn          net.Conn
}

func (c *Client) connect() {
	if c.Conn != nil {
		c.Close()
	}

	var err error
	fmt.Printf("Connecting to %s...\n", server)
	c.Conn, err = net.Dial("tcp", server+":"+port)
	if err != nil {
		fmt.Printf("Could not connect to %s, retrying in 5 seconds...\n", server)
		time.Sleep(time.Second * 5)
		c.connect()
	}
	fmt.Printf("Successfully connected to %s\n", server)
}

func (c *Client) login(pass, channel string) {
	fmt.Fprintf(c.Conn, "PASS %s\n", pass)
	fmt.Fprintf(c.Conn, "NICK %s\n", c.Nick)

	fmt.Fprintf(c.Conn, "JOIN %s\n", channel)
	fmt.Printf("Joined %s\n\n", channel)

	// Getting tags with the messages
	fmt.Fprintf(c.Conn, "CAP REQ :twitch.tv/tags\n")
}

// Close closes the clients connection
func (c *Client) Close() {
	c.Conn.Close()
	fmt.Printf("Closed the connection to %s\n", server)
}

// HandleChat handles incomming chat messages
func (c *Client) HandleChat() {
	proto := textproto.NewReader(bufio.NewReader(c.Conn))

	for {
		line, err := proto.ReadLine()
		if err != nil {
			log.Fatalf("err: %s", err)
			break
		}

		if strings.Contains(line, "PRIVMSG") {
			message := parser.ParseMessage(line, c.CommandPrefix)
			c.DisplayMessage(message)

			if message.IsCommand {
				parsedCmd := parser.ParseCommand(message, c.CommandPrefix)
				msg := command.HandleCommand(parsedCmd)
				c.SendMessage(msg, c.Channel)
			}
		}

		if strings.Contains(line, "PING :tmi.twitch.tv") {
			fmt.Fprintf(c.Conn, "PONG :tmi.twitch.tv\n")
		}
	}
}

// DisplayMessage displays incomming messages to the terminal
func (c *Client) DisplayMessage(msg *message.Message) {
	fmt.Printf("#%s %s: %s\n", msg.Channel, msg.User.Name, msg.Content)
}

// SendMessage sends message to chat
func (c *Client) SendMessage(msg, channel string) {
	if msg != "" {
		fmt.Fprintf(c.Conn, "PRIVMSG "+channel+" :"+msg+"\n")
	}
}

// NewClient, function generating new client
func NewClient(nick, pass, channel, prefix string) *Client {
	c := Client{
		Nick:          nick,
		Channel:       channel,
		CommandPrefix: prefix,
		Conn:          nil,
	}

	c.connect()
	c.login(pass, channel)

	return &c
}

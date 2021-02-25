package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"
	"time"
)

// Client object holding info about the connection
type Client struct {
	server  string
	port    string
	nick    string
	oAuth   string
	channel string
	conn    net.Conn
}

// NewClient, function generating new client
func NewClient(nick string, oAuth string, channel string) *Client {
	return &Client{
		server:  "irc.chat.twitch.tv",
		port:    "6667",
		nick:    nick,
		oAuth:   oAuth,
		channel: channel,
		conn:    nil,
	}
}

// Connect connects the client to the channel
func (c *Client) Connect() {
	if c.conn != nil {
		c.Close()
	}

	var err error
	fmt.Printf("Connecting to %s...\n", c.server)
	c.conn, err = net.Dial("tcp", c.server+":"+c.port)
	if err != nil {
		fmt.Printf("Could not connect to %s, retrying in 5 seconds...\n", c.server)
		time.Sleep(time.Second * 5)
		c.Connect()
	}
	fmt.Printf("Successfully connected to %s\n", c.server)
}

// Close closes the clients connection
func (c *Client) Close() {
	c.conn.Close()
	fmt.Printf("Closed the connection to %s\n", c.server)
}

// Login logs the client into the server
func (c *Client) Login() {
	fmt.Fprintf(c.conn, "PASS %s\n", c.oAuth)
	fmt.Fprintf(c.conn, "NICK %s\n", c.nick)

	fmt.Fprintf(c.conn, "JOIN %s\n", c.channel)
	fmt.Printf("Joined %s\n\n", c.channel)

	// Getting tags with the messages
	fmt.Fprintf(c.conn, "CAP REQ :twitch.tv/tags\n")
}

// HandleChat handles incomming chat messages
func (c *Client) HandleChat() {
	proto := textproto.NewReader(bufio.NewReader(c.conn))

	for {
		line, err := proto.ReadLine()
		if err != nil {
			log.Fatalln(err)
			break
		}

		if strings.Contains(line, "PRIVMSG") {
			message := ParseMessage(line, c.channel)
			c.DisplayMessage(message)

			if message.isCommand {
				parsedCommand := ParseMessageToCommand(message)
				HandleCommand(c, parsedCommand)
			}
		}

		if strings.Contains(line, "PING :tmi.twitch.tv") {
			fmt.Fprintf(c.conn, "PONG :tmi.twitch.tv\n")
		}
	}
}

// DisplayMessage displays incomming messages to the terminal
func (c *Client) DisplayMessage(msg *Message) {
	fmt.Printf("%s %s: %s\n", msg.channel, msg.user.displayname, msg.content)
}

// SendMessage sends message to chat
func (c *Client) SendMessage(msg string) {
	if msg != "" {
		fmt.Fprintf(c.conn, "PRIVMSG "+c.channel+" :"+msg+"\n")
	}
}

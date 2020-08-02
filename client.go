package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"strings"
	"time"
)

/*
Client -> object holding info about the connection
*/
type Client struct {
	server  string
	port    string
	nick    string
	oAuth   string
	channel string
	conn    net.Conn
}

/*
NewClient -> generates a new Client object
*/
func NewClient(nick string, oAuth string, channel string) *Client {
	return &Client{
		server:  "irc.chat.twitch.tv",
		port:    "6667",
		nick:    nick,
		oAuth:   oAuth,
		channel: channel,
		conn:    nil}
}

/*
Connect -> creates a connection between the client and the server
*/
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

/*
Close -> closes the clients connection to the server
*/
func (c *Client) Close() {
	c.conn.Close()
	fmt.Printf("Closed the connection to %s\n", c.server)
}

/*
Login -> logs the client into the channels
*/
func (c *Client) Login() {
	fmt.Fprintf(c.conn, "PASS %s\n", c.oAuth)
	fmt.Fprintf(c.conn, "NICK %s\n", c.nick)

	fmt.Fprintf(c.conn, "JOIN %s\n", c.channel)
	fmt.Printf("Joined %s\n\n", c.channel)

	// Getting tags with the messages
	fmt.Fprintf(c.conn, "CAP REQ :twitch.tv/tags\n")
}

/*
HandleChat -> handles incomming chats from the server
*/
func (c *Client) HandleChat() {
	proto := textproto.NewReader(bufio.NewReader(c.conn))

	for {
		line, err := proto.ReadLine()
		if err != nil {
			break
		}

		if strings.Contains(line, "PRIVMSG") {
			message := ParseMessage(line, c.channel)
			c.DisplayMessage(message)

			if message.isCommand {
				cmd := ParseCommand(message)
				cmd.Exec(c)
			}
		}

		if strings.Contains(line, "PING") {
			fmt.Fprintf(c.conn, "PONG :tmi.twitch.tv\n")
		}
	}
}

/*
DisplayMessage -> displays the formatted message in the console
*/
func (c *Client) DisplayMessage(msg *Message) {
	fmt.Printf("%s %s: %s\n", msg.channel, msg.user.displayname, msg.content)
}

/*
SendMessage -> sends a message to the channel
*/
func (c *Client) SendMessage(msg string) {
	if msg != "" {
		fmt.Fprintf(c.conn, "PRIVMSG "+c.channel+" :"+msg+"\n")
	}
}

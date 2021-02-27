package client

import (
	"fmt"
	"net"
	"time"

	"github.com/anders-14/bot_anders14_/pkg/command"
	"github.com/anders-14/bot_anders14_/pkg/irc"
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
	Channels      []string
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

func (c *Client) login(pass string) {
	fmt.Fprintf(c.Conn, "PASS %s\n", pass)
	fmt.Fprintf(c.Conn, "NICK %s\n", c.Nick)

	// Getting tags with the messages
	fmt.Fprintf(c.Conn, "CAP REQ :twitch.tv/tags\n")
}

func (c *Client) join(channel string) {
	fmt.Fprintf(c.Conn, "JOIN #%s\n", channel)
	fmt.Printf("Joined #%s\n", channel)
}

// Close closes the clients connection
func (c *Client) Close() {
	c.Conn.Close()
	fmt.Printf("Closed the connection to %s\n", server)
}

// HandleChat handles incomming chat messages
func (c *Client) HandleChat() {
	// Channels to communicate between go routines
	rawMessageChan := make(chan string, 100)
	messageChan := make(chan *message.Message, 100)
	commandChan := make(chan *message.Command, 100)
	pingChan := make(chan string)

	// Read and parse messages
	go irc.Read(c.Conn, rawMessageChan)
	go parser.Parse(rawMessageChan, messageChan, commandChan, pingChan, c.CommandPrefix)

	for {
		select {
		case cmd := <-commandChan:
			res := command.HandleCommand(cmd)
			if err := irc.Send(c.Conn, res, cmd.Channel); err != nil {
				fmt.Printf("err: %s", err)
			}
		case msg := <-messageChan:
			c.DisplayMessage(msg)
		case <-pingChan:
			irc.Pong(c.Conn)
		}
	}
}

// DisplayMessage displays incomming messages to the terminal
func (c *Client) DisplayMessage(msg *message.Message) {
	fmt.Printf("#%s %s: %s\n", msg.Channel, msg.User.Name, msg.Content)
}

// NewClient, function generating new client
func NewClient(nick, pass string, channels []string, prefix string) *Client {
	c := Client{
		Nick:          nick,
		Channels:      channels,
		CommandPrefix: prefix,
		Conn:          nil,
	}

	c.connect()
	c.login(pass)

	for _, channel := range channels {
		c.join(channel)
	}

	return &c
}

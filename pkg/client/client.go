package client

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"
	"time"

	"github.com/anders-14/bot_anders14_/pkg/parser"
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

func (c *Client) connect() {
	if c.conn != nil {
		c.Close()
	}

	var err error
	fmt.Printf("Connecting to %s...\n", c.server)
	c.conn, err = net.Dial("tcp", c.server+":"+c.port)
	if err != nil {
		fmt.Printf("Could not connect to %s, retrying in 5 seconds...\n", c.server)
		time.Sleep(time.Second * 5)
		c.connect()
	}
	fmt.Printf("Successfully connected to %s\n", c.server)
}

func (c *Client) login() {
	fmt.Fprintf(c.conn, "PASS %s\n", c.oAuth)
	fmt.Fprintf(c.conn, "NICK %s\n", c.nick)

	fmt.Fprintf(c.conn, "JOIN %s\n", c.channel)
	fmt.Printf("Joined %s\n\n", c.channel)

	// Getting tags with the messages
	fmt.Fprintf(c.conn, "CAP REQ :twitch.tv/tags\n")
}

// Close closes the clients connection
func (c *Client) Close() {
	c.conn.Close()
	fmt.Printf("Closed the connection to %s\n", c.server)
}

// HandleChat handles incomming chat messages
func (c *Client) HandleChat(cmdPrefix string) {
	proto := textproto.NewReader(bufio.NewReader(c.conn))

	for {
		line, err := proto.ReadLine()
		if err != nil {
			log.Fatalf("err: %s", err)
			break
		}

		if strings.Contains(line, "PRIVMSG") {
			message := parser.ParseMessage(line, cmdPrefix)
			c.DisplayMessage(message)

			if message.IsCommand {
				// parsedCommand := ParseMessageToCommand(message)
				// HandleCommand(c, parsedCommand)
				fmt.Println("its a command")
			}
		}

		if strings.Contains(line, "PING :tmi.twitch.tv") {
			fmt.Fprintf(c.conn, "PONG :tmi.twitch.tv\n")
		}
	}
}

// DisplayMessage displays incomming messages to the terminal
func (c *Client) DisplayMessage(msg *parser.Message) {
	fmt.Printf("#%s %s: %s\n", msg.Channel, msg.User.Name, msg.Content)
}

// SendMessage sends message to chat
func (c *Client) SendMessage(msg string) {
	if msg != "" {
		fmt.Fprintf(c.conn, "PRIVMSG "+c.channel+" :"+msg+"\n")
	}
}

// NewClient, function generating new client
func NewClient(nick string, oAuth string, channel string) *Client {
	c := Client{
		server:  "irc.chat.twitch.tv",
		port:    "6667",
		nick:    nick,
		oAuth:   oAuth,
		channel: channel,
		conn:    nil,
	}

	c.connect()
	c.login()

	return &c
}

package parser

import (
	"log"
	"regexp"
	"strings"

	"github.com/anders-14/bot_anders14_/pkg/message"
)

var (
	contentRegex = regexp.MustCompile(`@([^\s]+)\s:([^!]+)![^#]+#([^\s]+)\s:(.+)`)
)

func isPrivateMessage(line string) bool {
	return strings.Contains(line, "PRIVMSG")
}

func isPingMessage(line string) bool {
	return strings.Contains(line, "PING :tmi.twitch.tv")
}

func parseList(str, listSep, valSep string) map[string]string {
	itemMap := make(map[string]string)

	items := strings.Split(str, listSep)

	for _, item := range items {
		parts := strings.Split(item, valSep)
		itemMap[parts[0]] = parts[1]
	}

	return itemMap
}

func parseTags(tagString string) map[string]string {
	return parseList(tagString, ";", "=")
}

func parseBadges(badgeString string) map[string]string {
	return parseList(badgeString, ",", "/")
}

func ParseCommand(msg *message.Message, cmdPrefix string) *message.Command {
	prefixLen := len(cmdPrefix)
	splitMessage := strings.Split(msg.Content, " ")
	name := strings.ToLower(splitMessage[0][prefixLen:])
	args := splitMessage[1:]

	return &message.Command{
		Name:    name,
		Args:    args,
		User:    msg.User,
		Channel: msg.Channel,
	}
}

func parseMessage(line, cmdPrfix string) *message.Message {
	matches := contentRegex.FindAllStringSubmatch(line, 5)[0]

	tagString := matches[1]
	username := matches[2]
	channelname := matches[3]
	content := matches[4]

	tags := parseTags(tagString)
	badgeString, ok := tags["badges"]
	if !ok {
		log.Fatalf("err: no badges in tags")
	}
	badges := parseBadges(badgeString)

	_, isBroad := badges["broadcaster"]
	_, isMod := badges["moderator"]
	_, isVip := badges["vip"]
	_, isSub := badges["subscriber"]

	msg := message.Message{
		Content: content,
		User: message.User{
			ID:            tags["user-id"],
			Name:          username,
			Badges:        badges,
			Color:         tags["color"],
			IsBroadcaster: isBroad,
			IsModerator:   isMod,
			IsVip:         isVip,
			IsSubscriber:  isSub,
		},
		Channel:   channelname,
		IsCommand: strings.HasPrefix(content, cmdPrfix),
	}
	return &msg
}

func Parse(raw <-chan string, msgs chan *message.Message, cmds chan *message.Command, pings chan string, cmdPrefix string) {
	for {
		line := <-raw
		if isPrivateMessage(line) {
			msg := parseMessage(line, cmdPrefix)
			msgs <- msg
			if msg.IsCommand {
				cmd := ParseCommand(msg, cmdPrefix)
				cmds <- cmd
			}
		} else if isPingMessage(line) {
			pings <- line
		}
	}
}

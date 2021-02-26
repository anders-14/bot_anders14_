package parser

import (
	// "fmt"
	"log"
	"regexp"
	"strings"
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

// ParseMessage parses irc message into message object
func ParseMessage(line, cmdPrfix string) *Message {
	matches := contentRegex.FindAllStringSubmatch(line, 5)[0]

	tagString := matches[1]
	channelname := matches[2]
	username := matches[3]
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

	msg := Message{
		Content: content,
		User: User{
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

// @badge-info=;badges=broadcaster/1;client-nonce=cd9e3395f945b2d5c92680572bc64a41;color=#FF7F50;display-name=anders14_;emotes=;flags=;id=12297ad5-e164-45da-a9a5-dc46446aadfe;mod=0;room-id=207792212;subscriber=0;tmi-sent-ts=1614326479328;turbo=0;user-id=207792212;user-type=
// :anders14_!anders14_@anders14_.tmi.twitch.tv PRIVMSG #anders14_ :test

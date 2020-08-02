package main

import (
	"regexp"
	"strconv"
	"strings"
)

/*
Message -> object holding info about a message
*/
type Message struct {
	channel   string
	content   string
	isCommand bool
	user      User
}

/*
User -> object holding info about a user
*/
type User struct {
	badges        []Badge
	isBroadcaster bool
	color         string
	displayname   string
	id            string
	isModerator   bool
	isSubscriber  bool
	username      string
}

/*
Badge -> object holding info about a badge
*/
type Badge struct {
	name  string
	value string
}

/*
ParseMessage -> parsing the message from twitch
and returning the parsed message object
*/
func ParseMessage(message string, channel string) *Message {
	// Gets the message contents
	contentRegex := regexp.MustCompile(`PRIVMSG\s#\w+\s:(.+)`)
	content := contentRegex.FindStringSubmatch(message)[1]

	// Is the message a command
	isCommand := strings.HasPrefix(content, "!")

	// Gets the user-tags
	tagsRegex := regexp.MustCompile(`(@.+):\w+!`)
	tagsString := tagsRegex.FindStringSubmatch(message)[1]
	tagsList := strings.Split(tagsString, ";")

	// The user object from tagsList
	user := UserFromTags(tagsList)

	return &Message{
		channel:   channel,
		content:   content,
		isCommand: isCommand,
		user:      *user}
}

/*
UserFromTags -> gets a user object from tags
*/
func UserFromTags(tags []string) *User {
	tagMap := make(map[string]string)

	for _, v := range tags {
		valuePair := strings.Split(v, "=")
		key := valuePair[0]
		value := valuePair[1]

		tagMap[key] = value
	}

	// The users badges
	var badges []Badge
	rawBadges := strings.Split(tagMap["badges"], ",")
	for _, v := range rawBadges {
		badgeInfo := strings.Split(v, "/")
		if len(badgeInfo) > 1 {
			badge := &Badge{
				name:  badgeInfo[0],
				value: badgeInfo[1]}
			badges = append(badges, *badge)
		}
	}

	// Is the user also the broadcaster
	broadcasterStatus := false
	for _, v := range badges {
		if v.name == "broadcaster" {
			broadcasterStatus = true
		}
	}

	// Is the user a moderator
	modStatus, err := strconv.ParseBool(tagMap["mod"])
	if err != nil {
		modStatus = false
	}

	// Is the user a subscriber
	subStatus, err := strconv.ParseBool(tagMap["subscriber"])
	if err != nil {
		subStatus = false
	}

	return &User{
		badges:        badges,
		isBroadcaster: broadcasterStatus,
		color:         tagMap["color"],
		displayname:   tagMap["display-name"],
		id:            tagMap["id"],
		isModerator:   modStatus,
		isSubscriber:  subStatus,
		username:      strings.ToLower(tagMap["display-name"])}
}

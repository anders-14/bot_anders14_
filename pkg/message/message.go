package message

// Message is parsed irc messages
type Message struct {
	Content   string
	User      User
	Channel   string
	IsCommand bool
}

// Command is a parsed bot command
type Command struct {
	Name    string
	Args    []string
	User    User
	Channel string
}

// User parsed from the irc message tags
type User struct {
	ID            string
	Name          string
	Badges        map[string]string
	Color         string
	IsBroadcaster bool
	IsModerator   bool
	IsVip         bool
	IsSubscriber  bool
}

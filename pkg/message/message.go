package message

type Message struct {
	Content   string
	User      User
	Channel   string
	IsCommand bool
}

type Command struct {
  Name string
  Args []string
  User User
  Channel string
}

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

package main

func main() {
	client := NewClient(
		"yourBotName",
		"oauth:yourToken",
		"#twitchChannel",
	)

	client.Connect()
	defer client.Close()
	client.Login()

	client.HandleChat()
}

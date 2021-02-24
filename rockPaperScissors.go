package main

import (
	"math/rand"
	"strings"
)

var moves [3]string = [3]string{"rock", "paper", "scissors"}

// PlayRPS -> plays rps against the user
func PlayRPS(usermove string) string {
	usermove = strings.ToLower(usermove)
	if invalidUserMove(usermove) {
		return ""
	}

	botmove := getRandomMove()

	if checkDraw(usermove, botmove) {
		return "I pick " + botmove + ", it's a draw"
	}

	if userWins(usermove, botmove) {
		return "I pick " + botmove + ", you win"
	}

	return "I pick " + botmove + ", you lose"
}

func invalidUserMove(usermove string) bool {
	return usermove != "rock" && usermove != "paper" && usermove != "scissors"
}

func getRandomMove() string {
	return moves[rand.Intn(3)]
}

func checkDraw(usermove string, botmove string) bool {
	return usermove == botmove
}

func userWins(usermove string, botmove string) bool {
	if usermove == "rock" {
		return botmove == "scissors"
	}
	if usermove == "paper" {
		return botmove == "rock"
	}
	if usermove == "scissors" {
		return botmove == "paper"
	}

	return false
}

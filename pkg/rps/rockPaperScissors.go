package rps

import (
	"fmt"
	"math/rand"
	"strings"
)

var moves [3]string = [3]string{"rock", "paper", "scissors"}

// check if the move is valid
func validateMove(move string) bool {
	for _, v := range moves {
		if move == v {
			return true
		}
	}
	return false
}

// get random move
func randomMove() string {
  return moves[rand.Intn(3)]
}

// checks if the moves are the same
func checkDraw(move1, move2 string) bool {
  return move1 == move2
}

func userWon(usermove, botmove string) bool {
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

// Play, given a usermove and name plays a round of rps
func Play(usermove, username string) string {
  loweredMove := strings.ToLower(usermove)

	if !validateMove(loweredMove) {
    return fmt.Sprintf("%s is not a valid rps move, @%s", usermove, username)
	}

  botmove := randomMove()

  if checkDraw(loweredMove, botmove) {
    return fmt.Sprintf("I pick %s, its a draw @%s", botmove, username)
  }

  if userWon(loweredMove, botmove) {
    return fmt.Sprintf("I pick %s, you win @%s", botmove, username)
  }

  return fmt.Sprintf("I pick %s, you lose @%s", botmove, username)
}

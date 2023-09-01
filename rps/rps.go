package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 // Stone. Defeat scissors (scissors +1) % 3 = 0
	PAPER    = 1 // Paper. Defeat stone (stone + 1) % 3 = 1
	SCISSORS = 2 // Scissor. Defeat paper. (paper + 1) % 3 = 3
)

// Structure to give results for each round

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json: "computer_choice"`
	RoundResult       string `json: "round_result"`
	ComputerChoiceInt int    `json: "computer_choise_int"`
	ComputerScore     string `json: "computer_score"`
	PlayerScore       string `json: "player_score"`
}

// Message for when you win
var winMessages = []string{
	"Well done!",
	"Good job!",
	"You should buy a lottery ticket",
}

// Mesage for when you lose
var loseMessages = []string{
	"What a pity!",
	"Try again!",
	"Today just isn't your day.",
}

// Draw message
var drawMessages = []string{
	"Great minds think alike.",
	"Oh oh. Try again.",
	"Nobody wins, buy you can try again.",
}

// Scoring variables
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	// Message depending on what the computer chose
	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "The computer chose STONE"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "The computer chose PAPER"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "The computer chose scissors"
	}

	// generate a random number from 0-2, which we us to choose a random message
	messageInt := rand.Intn(3)

	// declare a var to contain the message
	var message string

	if playerValue == computerValue {
		roundResult = "It is a tie"
		// select message from drawMessages
		message = drawMessages[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "The player wins!"
		// select message from winMessages
		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "Computer wins!"
		// select message from loseMessages
		message = loseMessages[messageInt]
	}

	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}

package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK      = 0 //Piedra. Vence a las tijeras. (tijeras + 1) % 3 = 0
	PAPER     = 1 // Papel. Vence a la piedra. (piedra + 1) % 3 = 1
	SCCISSORS = 2 // Tijeras. Vencen al papel. (papel + 1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_client"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

var winMessages = []string{
	"Bien hecho!",
	"Buen trabajo!",
	"Deberias comprar un boleto de loteria",
}

var loseMessages = []string{
	"Que lastima!",
	"Intentalo de nuevo!",
	"Hoy simplemente no es tu dia!",
}

var drawMessagess = []string{
	"Las grandes mentes piensan igual.",
	"Oh oh. Intentalo de nuevo.",
	"Nadie gana, pero puedes intentarlo de nuevo",
}

var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PIEDRA"
	case PAPER:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio PAPEL"
	case SCCISSORS:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligio TIJERAS"
	}

	messageInt := rand.Intn(3)

	var message string

	if playerValue == computerValue {
		roundResult = "Es un empate"

		message = drawMessagess[messageInt]
	} else if playerValue == (computerValue+1)%3 {
		PlayerScore++
		roundResult = "El jugador gana"

		message = winMessages[messageInt]
	} else {
		ComputerScore++
		roundResult = "La computadora gana!"
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

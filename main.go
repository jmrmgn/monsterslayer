package main

import (
	"github.com/jmrmgn/monsterslayer/actions"
	"github.com/jmrmgn/monsterslayer/interaction"
)

var currentRound = 0

func main() {
	startGame()

	winner := "" // "Player" || "Monster" || ""

	for winner == "" {
		winner = executeRound()
	}

	endGame()
}

func startGame() {
	interaction.PrintGreeting()
}

func executeRound() string {
	currentRound++
	isSpecialRound := currentRound%3 == 0

	interaction.ShowAvailableActions(isSpecialRound)
	userChoice := interaction.GetPlayerChoice(isSpecialRound)

	if userChoice == "ATTACK" {
		actions.AttackMonster(false)
	} else if userChoice == "HEAL" {
		actions.HealPlayer()
	} else {
		actions.AttackMonster(isSpecialRound)
	}

	actions.AttackPlayer()

	playerHp, monsterHp := actions.GetHPAmounts()

	if playerHp <= 0 {
		return "Monster"
	} else if monsterHp <= 0 {
		return "Player"
	}

	return ""
}

func endGame() {}

package actions

import (
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)
var currentMonsterHP = MONSTER_HP
var currentPlayerHP = PLAYER_HP

func AttackMonster(isSpecialAttack bool) int {
	minAttackValue := PLAYER_ATTACK_MIN_DMG
	maxAttackValue := PLAYER_ATTACK_MAX_DMG

	if isSpecialAttack {
		minAttackValue = PLAYER_SPECIAL_ATTACK_MIN_DMG
		maxAttackValue = PLAYER_SPECIAL_ATTACK_MAX_DMG
	}

	dmgValue := generateRandBetween(minAttackValue, maxAttackValue)
	currentMonsterHP -= dmgValue

	return dmgValue
}

func HealPlayer() int {
	healValue := generateRandBetween(PLAYER_HEAL_MIN_VALUE, PLAYER_HEAL_MAX_VALUE)

	healthDiff := PLAYER_HP - currentPlayerHP

	if healthDiff >= healValue {
		currentPlayerHP += healValue
		return healValue
	} else {
		currentPlayerHP = PLAYER_HP
		return healthDiff
	}
}

func AttackPlayer() int {
	dmgValue := generateRandBetween(MONSTER_ATTACK_MIN_DMG, MONSTER_ATTACK_MAX_DMG)
	currentPlayerHP -= dmgValue

	return dmgValue
}

func GetHPAmounts() (int, int) {
	return currentPlayerHP, currentMonsterHP
}

func generateRandBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}

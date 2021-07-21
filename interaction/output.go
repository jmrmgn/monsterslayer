package interaction

import "fmt"

type RoundData struct {
	Action           string
	PlayerAttackDmg  int
	PlayerHealValue  int
	MonsterAttackDmg int
	PlayerHP         int
	MonsterHP        int
}

func PrintGreeting() {
	fmt.Println("MONSTER SLAYER")
	fmt.Println("Starting a new game...")
	fmt.Println("Good luck!")
}

func ShowAvailableActions(specialAttackIsAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("--------------------------")

	fmt.Println("(1) Attack Monster")
	fmt.Println("(2) Heal")

	if specialAttackIsAvailable {
		fmt.Println("(3) Special Attack")
	}
}

func PrintRoundStatistics(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed a special attack against monster for %v damage.\n", roundData.PlayerAttackDmg)
	} else if roundData.Action == "HEAL" {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackDmg)
	fmt.Printf("Player HP: %v\n", roundData.PlayerHP)
	fmt.Printf("Monster HP: %v\n", roundData.MonsterHP)
}

func DeclareWinner(winner string) {
	fmt.Println("--------------------------")
	fmt.Println("GAME OVER!")
	fmt.Println("--------------------------")
	fmt.Printf("%v won \n", winner)
}

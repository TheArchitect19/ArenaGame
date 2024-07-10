package main

import (
	"ArenaGame/modules"
	"ArenaGame/utils"
	"fmt"
)

func inputNewPlayerDetails() (string, int, int, int) {
	name := utils.InputStringFromUser("Enter the player's name: ")
	health := utils.InputIntegerFromUser(fmt.Sprintf("Enter %s's health: ", name))
	attack := utils.InputIntegerFromUser(fmt.Sprintf("Enter %s's attack: ", name))
	strength := utils.InputIntegerFromUser(fmt.Sprintf("Enter %s's strength: ", name))
	return name, health, attack, strength
}

func mainExecuter() {
	arena := modules.NewArena()
	for {
		arena.DisplayPlayers()
		fmt.Println("Options: \n\t1> Add new player\n\t2> Battle\n\t3> End game\n")
		option := utils.InputIntegerFromUser("Enter your choice (integer): ")
		switch option {
		case 1:
			name, health, attack, strength := inputNewPlayerDetails()
			arena.AddPlayer(name, health, strength, attack)
		case 2:
			if arena.GetPlayerCount() < 2 {
				fmt.Println("There should be at least two players in the Arena.\nPlease add more players.\n")
			} else {
				idFirst := utils.InputIntegerFromUser("Enter the first player's id: ")
				idSecond := utils.InputIntegerFromUser("Enter the second player's id: ")
				arena.Battle(idFirst, idSecond)
			}
		default:
			fmt.Println("Bye Bye...\n")
			return
		}
		fmt.Println("\n____________________________________________________________________________________________________\n\n")
	}
}

func main() {
	mainExecuter()
}

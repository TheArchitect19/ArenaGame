package modules

import (
	"fmt"
	"ArenaGame/utils"
)


// Arena represents the arena where players battle.
type Arena struct {
	TotalPlayers int
	Players      map[int]*Player
}

// NewArena creates a new arena.
func NewArena() *Arena {
	return &Arena{
		TotalPlayers: 0,
		Players:      make(map[int]*Player),
	}
}

// IsPresent checks if a player is present in the arena.
func (a *Arena) IsPresent(id int) bool {
	_, exists := a.Players[id]
	return exists
}

// GetPlayerCount returns the count of players in the arena.
func (a *Arena) GetPlayerCount() int {
	return len(a.Players)
}

// AddPlayer adds a new player to the arena.
func (a *Arena) AddPlayer(name string, health, strength, attack int) int {
	if health <= 0 {
		fmt.Println("Health should be a positive integer.")
		return -1
	}
	if strength <= 0 {
		fmt.Println("Strength should be a positive integer.")
		return -1
	}
	if attack <= 0 {
		fmt.Println("Attack should be a positive integer.")
		return -1
	}

	id := a.TotalPlayers
	newPlayer := NewPlayer(id, name, health, strength, attack)
	a.Players[id] = newPlayer
	a.TotalPlayers++
	return id
}

// DeletePlayer removes a player from the arena.
func (a *Arena) DeletePlayer(id int) {
	if player, exists := a.Players[id]; exists {
		fmt.Printf("%s has been un-alived...\n", player.Name)
		delete(a.Players, id)
	} else {
		fmt.Printf("No player with id = %d exists.\n", id)
	}
}

// DisplayPlayers displays the list of players in the arena.
func (a *Arena) DisplayPlayers() {
	fmt.Println("|\tId\t|\tName\t|\tHealth\t|\tStrength\t|\tAttack\t|")
	for id, player := range a.Players {
		fmt.Printf("|\t%d\t|\t%s\t|\t%d\t|\t%d\t|\t%d\t|\n", id, player.Name, player.Health, player.Strength, player.Attack)
	}
	fmt.Println()
}

// Battle initiates a battle between two players.
func (a *Arena) Battle(idFirst, idSecond int) map[string]int {
	if idFirst == idSecond {
		fmt.Println("Id's can not be the same for both the players.\n")
		return map[string]int{}
	}
	if !a.IsPresent(idFirst) {
		fmt.Printf("No player with id = %d exists.\n", idFirst)
		return map[string]int{}
	}
	if !a.IsPresent(idSecond) {
		fmt.Printf("No player with id = %d exists.\n", idSecond)
		return map[string]int{}
	}

	attacker := a.Players[idFirst]
	defender := a.Players[idSecond]
	fmt.Printf("\n____________%s v/S %s____________\n", attacker.Name, defender.Name)

	if defender.Health < attacker.Health {
		attacker, defender = defender, attacker
	}

	for defender.Health > 0 {
		attackingPower := attacker.Attack * utils.RollDice()
		defendingPower := defender.Strength * utils.RollDice()

		fmt.Printf("%s hits %s with power = %d\n", attacker.Name, defender.Name, attackingPower)
		fmt.Printf("%s defends with power = %d\n", defender.Name, defendingPower)

		if attackingPower > defendingPower {
			defender.Health -= (attackingPower - defendingPower)
			if defender.Health < 0 {
				defender.Health = 0
			}
		}

		fmt.Printf("%s's health: %d\n", defender.Name, defender.Health)

		if defender.Health > 0 {
			attacker, defender = defender, attacker
		}
	}

	res := map[string]int{"winner": attacker.ID, "loser": defender.ID}
	fmt.Printf("%s has emerged victorious!!!\n", attacker.Name)
	a.DeletePlayer(defender.ID)

	return res
}
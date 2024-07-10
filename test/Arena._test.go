package test

import (
	"testing"
	"ArenaGame/modules"
)

func TestAddPlayer(t *testing.T) {
	arena := modules.NewArena() 

	t.Run("Player's health can't be negative or zero", func(t *testing.T) {
		id := arena.AddPlayer("A", -2, 200, 100)
		if id != -1 {
			t.Errorf("Expected id to be -1, got %d", id)
		}

		id = arena.AddPlayer("A", 0, 200, 100)
		if id != -1 {
			t.Errorf("Expected id to be -1, got %d", id)
		}
	})

	t.Run("Player's strength can't be negative or zero", func(t *testing.T) {
		id := arena.AddPlayer("A", 2, -200, 100)
		if id != -1 {
			t.Errorf("Expected id to be -1, got %d", id)
		}

		id = arena.AddPlayer("A", 10, 0, 100)
		if id != -1 {
			t.Errorf("Expected id to be -1, got %d", id)
		}
	})

	t.Run("Player's attack can't be negative or zero", func(t *testing.T) {
		id := arena.AddPlayer("A", 2, 200, -100)
		if id != -1 {
			t.Errorf("Expected id to be -1, got %d", id)
		}

		id = arena.AddPlayer("A", 10, 120, 0)
		if id != -1 {
			t.Errorf("Expected id to be -1, got %d", id)
		}
	})

	t.Run("Newly added player to be present in the Arena", func(t *testing.T) {
		id := arena.AddPlayer("A", 100, 200, 100)
		if !arena.IsPresent(id) {
			t.Errorf("Expected player with id %d to be present", id)
		}
	})

	t.Run("Player count should increase after addition of a new player", func(t *testing.T) {
		oldPlayerCount := arena.GetPlayerCount()
		arena.AddPlayer("A", 100, 200, 100)
		newPlayerCount := arena.GetPlayerCount()

		if newPlayerCount != oldPlayerCount+1 {
			t.Errorf("Expected player count to be %d, got %d", oldPlayerCount+1, newPlayerCount)
		}
	})
}

func TestDeletePlayer(t *testing.T) {
	arena := modules.NewArena()

	t.Run("Deleted player should not be present in the Arena", func(t *testing.T) {
		id := arena.AddPlayer("A", 100, 200, 100)
		arena.DeletePlayer(id)
		if arena.IsPresent(id) {
			t.Errorf("Expected player with id %d to be absent", id)
		}
	})

	t.Run("Player count should decrease after deletion of a player", func(t *testing.T) {
		id := arena.AddPlayer("A", 100, 200, 100)
		oldPlayerCount := arena.GetPlayerCount()
		arena.DeletePlayer(id)
		newPlayerCount := arena.GetPlayerCount()

		if newPlayerCount != oldPlayerCount-1 {
			t.Errorf("Expected player count to be %d, got %d", oldPlayerCount-1, newPlayerCount)
		}
	})

	t.Run("Player not in the Arena should not be deleted", func(t *testing.T) {
		id := arena.AddPlayer("A", 100, 200, 100)
		oldPlayerCount := arena.GetPlayerCount()
		arena.DeletePlayer(id + 123)
		newPlayerCount := arena.GetPlayerCount()

		if newPlayerCount != oldPlayerCount {
			t.Errorf("Expected player count to be %d, got %d", oldPlayerCount, newPlayerCount)
		}
	})
}

func TestBattle(t *testing.T) {
	arena := modules.NewArena()

	t.Run("Battle with empty arena", func(t *testing.T) {
		result := arena.Battle(0, 1)
		if len(result) != 0 {
			t.Errorf("Expected result to be empty, got %v", result)
		}
	})

	t.Run("Players have the same id", func(t *testing.T) {
		arena.AddPlayer("A", 100, 200, 100)
		arena.AddPlayer("B", 200, 300, 100)

		result := arena.Battle(0, 0)
		if len(result) != 0 {
			t.Errorf("Expected result to be empty, got %v", result)
		}
	})

	t.Run("One of the Player's id does not exist", func(t *testing.T) {
		arena.AddPlayer("A", 100, 200, 100)
		arena.AddPlayer("B", 200, 300, 100)

		result := arena.Battle(0, 10)
		if len(result) != 0 {
			t.Errorf("Expected result to be empty, got %v", result)
		}
	})

	t.Run("Normal battle", func(t *testing.T) {
		arena.AddPlayer("A", 100, 200, 100)
		arena.AddPlayer("B", 200, 300, 100)

		possibleOutcomes := []map[string]int{{"winner": 0, "loser": 1}, {"winner": 1, "loser": 0}}

		result := arena.Battle(0, 1)

		found := false
		for _, outcome := range possibleOutcomes {
			if outcome["winner"] == result["winner"] && outcome["loser"] == result["loser"] {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected result to be one of %v, got %v", possibleOutcomes, result)
		}

		winner, loser := result["winner"], result["loser"]

		if !arena.IsPresent(winner) {
			t.Errorf("Expected winner with id %d to be present", winner)
		}
		if arena.IsPresent(loser) {
			t.Errorf("Expected loser with id %d to be absent", loser)
		}
	})
}

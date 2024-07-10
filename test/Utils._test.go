package test

import (
	"math/rand"
	"testing"
	"time"
)

// rollDice simulates rolling a six-sided die.
func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	return rand.Intn(max-min+1) + min
}

func TestRollDice(t *testing.T) {
	t.Run("returns a number between 1 and 6", func(t *testing.T) {
		result := rollDice()
		if result < 1 || result > 6 {
			t.Errorf("Expected result to be between 1 and 6, got %d", result)
		}
	})

	t.Run("returns an integer", func(t *testing.T) {
		result := rollDice()
		if result != int(result) {
			t.Errorf("Expected result to be an integer, got %v", result)
		}
	})
}

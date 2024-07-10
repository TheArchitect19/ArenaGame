package utils

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// rollDice simulates rolling a six-sided die.
func RollDice() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 6
	return rand.Intn(max-min+1) + min
}

// inputStringFromUser prompts the user for input and returns the entered string.
func InputStringFromUser(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// inputIntegerFromUser prompts the user for input and returns the entered integer.
func InputIntegerFromUser(prompt string) int {
	for {
		input := InputStringFromUser(prompt)
		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}
		return value
	}
}

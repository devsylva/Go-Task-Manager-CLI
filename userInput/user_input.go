package userInput

import (
	"fmt"
)

// readstring reads a string from the user
func ReadString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
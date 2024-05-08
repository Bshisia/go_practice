package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// Check if the number of arguments is not equal to 2
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	// Extract the two strings from command-line arguments
	firstString := os.Args[1]
	secondString := os.Args[2]

	// Index to iterate over the second string
	index := 0

	// Iterate over the characters of the first string
	for _, char := range firstString {
		// Check if the character exists in the second string
		found := false
		for ; index < len(secondString); index++ {
			if rune(secondString[index]) == char {
				index++ // Move to the next character in the second string
			}
		}
		// If the character is not found in the second string, it's not possible to write the first string
		if !found {
			z01.PrintRune('\n')
			return
		}
	}

	// If all characters of the first string are found in the second string in the same order
	// then it's possible to write the first string
	for _, char := range firstString {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}

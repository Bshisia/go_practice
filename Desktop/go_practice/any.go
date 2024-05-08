package main

import (
	"fmt"
	"unicode"
)

func Any(f func(string) bool, str []string) bool {
	for _, value := range str {
		if f(value) {
			return true
		}
	}
	return false
}

func main() {
	hasNumeric := Any(func(s string) bool {
		for _, char := range s {
			if unicode.IsDigit(char) {
				return true
			}
		}
		return false
	}, []string{"Hello", "this", "4", "yo"})

	fmt.Println(hasNumeric) // Output: true (because "4" contains a numeric figure)
}

package main

import "fmt"

func AtoiBase(s string, base string) int {
	result := 0
	len := 0
	array := map[rune]bool{}
	for _, bass := range base {
		if array[bass] || bass == '-' || bass == '+' {
			return result
		}
		array[bass] = true
		len++
	}
	if len > 1 {
		for _, str := range s {
			count := 0
			if array[str] {
				for _, bas := range base {
					if bas == str {
						break
					}
					count++
				}
				result = result*10 + count
			}
		}

	}
	return result
}
func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("oui", "choumi"))
}

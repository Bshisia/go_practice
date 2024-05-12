package main

import "fmt"

func BasicAtoi(s string) int {
	result := 0
	Negative := false
	for i, str := range s {
		if str >= '0' && str <= '9' {
			result = result*10 + int(str-'0')
		} else if i == 0 && (str == '-' || str == '+') {
			if str == '-' {
				Negative = true
			}
		} else {
			return 0
		}
	}
	if Negative {
		result = result * -1
	}
	return result
}

func main() {
	char := "7D"
	charr := BasicAtoi(char)
	fmt.Println(charr)
}

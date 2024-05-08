package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		z01.PrintRune('0')
		return
	}
	n := atoi(args[1])
	if n <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	sum := 0
	for i := 0; i <= n; i++ {
		if IsPrime(i) {
			sum += i
		}
	}
	res := Itoa(sum)
	for _, digits := range res {
		z01.PrintRune(digits)
	}

	z01.PrintRune('\n')
}
func atoi(s string) int {
	result := 0
	for _, digits := range s {
		if digits >= '0' && digits <= '9' {
			result = result*10 + int(digits-'0')
		} else {
			return -1
		}
	}
	return result
}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var result string
	for n > 0 {
		digit := n % 10
		result = string('0'+digit) + result
		n /= 10
	}
	return result
}

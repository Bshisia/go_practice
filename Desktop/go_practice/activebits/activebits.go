package main

import "fmt"

func ActiveBits(n int) uint {
	count := uint(0)
	for n > 0 {
		count += uint(n & 1) // Count the rightmost bit
		n >>= 1              // Shift n to the right by 1
	}
	return count
}

func main() {
	nbits := ActiveBits(9)
	fmt.Println(nbits)
}

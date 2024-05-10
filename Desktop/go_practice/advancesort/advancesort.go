package main

import "fmt"

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if f(a[i], a[j]) > 0 {
				a[j], a[i] = a[i], a[j]
			}
		}
	}
}

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func main() {
	result := []string{"a", "4", "1", "b", "2", "B", "c", "C", "3"}
	AdvancedSortWordArr(result, Compare)
	fmt.Println(result)
}

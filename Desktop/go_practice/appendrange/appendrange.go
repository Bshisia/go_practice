package main

import "fmt"

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var result  []int
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}

func main() {
	appended := AppendRange(10, 5)
	fmt.Println(appended)
}

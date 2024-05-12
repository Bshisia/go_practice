package main

import "fmt"

func AppendRange(min, max int) []int {
	var res []int
	for i := min; i <= max; i++ {
		res = append(res, i)
	}
	return res
}

func main() {
	appended := AppendRange(10, 15)
	fmt.Println(appended)
}

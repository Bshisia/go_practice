package main

import "fmt"

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	var res []int
	for i := min; i < max; i++ {
		res = append(res, i)
	}
	return res
}

func main() {
	appended := AppendRange(10, 5)
	fmt.Println(appended)
}
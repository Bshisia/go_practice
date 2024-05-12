package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error in Reading file")
	}
	// fmt.Println(string(file))
	stringslice := strings.Fields(string(file))
	stringslice = HexBin(stringslice)
	stringslice = Format(stringslice)
	stringslice = Vowels(stringslice)
	stringslice = Quotations(stringslice)
	stringslice = Punctuate(stringslice)
	finalstring := strings.Join(stringslice, " ")
	os.WriteFile(os.Args[2], []byte(finalstring), 0644)

}

func HexBin(str []string) []string {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i], "(hex)") {
			inter, err := strconv.ParseInt(str[i-1], 16, 64)
			if err != nil {
				fmt.Println("Error in conversion")
			}
			str[i-1] = strconv.Itoa(int(inter))
			str = append(str[:i], str[i+1:]...)
		}
	}
	return str
}

func Format(str []string) []string {
	for i := 0; i < len(str); i++ {
		if strings.Contains(str[i], "(cap)") {
			str[i-1] = strings.ToUpper(str[i-1][:1]) + str[i-1][1:]
		}
		if strings.Contains(str[i], "(cap,") {
			num, err := strconv.Atoi(str[i+1][:len(str[i+1])-1])
			if err != nil {
				fmt.Println("Error, in cap conversion")
			}
			for j := i - num; j < i; j++ {
				str[j] = strings.ToUpper(str[j][:1]) + str[j][1:]
			}
			str = append(str[:i], str[i+2:]...)
		}
	}
	return str
}

func Punctuate(str []string) []string {
	puncs := []string{",", ".", "?", ",", "!", ";", ":"}
	for i := 0; i < len(str); i++ { // Start from 1 to avoid accessing index -1
		for _, puns := range puncs {
			if strings.HasPrefix(str[i], puns) {
				str[i-1] += puns
				str[i] = str[i][1:]   // Remove the punctuation mark
				return Punctuate(str) // Recursive call to check for more punctuation
			}
		}
	}
	return str
}

func Quotations(str []string) []string {
	count := 0
	quotes := "'"
	for i := 0; i < len(str); i++ {
		if str[i] == quotes && count == 0 {
			count++
			str[i+1] = quotes + str[i+1]
			str = append(str[:i], str[i+1:]...)
		} else if str[i] == quotes && count > 0 {
			str[i-1] = str[i-1] + quotes
			str = append(str[:i], str[i+1:]...)
		}
	}
	return str
}

func Vowels(str []string) []string {
	Vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	for i := 0; i < len(str); i++ {
		for _, vow := range Vowels {
			if strings.HasPrefix(str[i], vow) && str[i-1] == "a" {
				str[i-1] = "an"
			} else if strings.HasPrefix(str[i], vow) && str[i-1] == "A" {
				str[i-1] = "An"
			}
		}
	}
	return str
}

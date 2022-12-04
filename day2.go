package main

import (
	"fmt"
	"strings"
)

func day2(get_data string) {
	list_data := strings.Split(get_data, "\n")
	score := 0
	var dictionaryValues = make(map[string]int)
	dictionaryValues["A"] = 1
	dictionaryValues["B"] = 2
	dictionaryValues["C"] = 3
	dictionaryValues["X"] = 1
	dictionaryValues["Y"] = 2
	dictionaryValues["Z"] = 3

	for _, data := range list_data {
		tmp_data := strings.Split(data, " ")
		if data == "A X" || data == "B Y" || data == "C Z" {
			score += dictionaryValues[tmp_data[1]] + 3
		} else if data == "C X" || data == "A Y" || data == "B Z" {
			score += dictionaryValues[tmp_data[1]] + 6
		} else if data == "A Z" || data == "B X" || data == "C Y" {
			score += dictionaryValues[tmp_data[1]]
		}
	}
	fmt.Println(score)
}

package main

import (
	"fmt"
	"strings"
)

func day3(data_string string) {
	tmp_input := strings.ReplaceAll(data_string, "\r", "")
	rucksacks := strings.Split(tmp_input, "\n")
	total_result := 0

	//anonymous function to calc_priority
	calc_priority := func(priorities_maping map[rune]int, individual_rucksack string) int {
		comp1 := individual_rucksack[:len(individual_rucksack)/2]
		comp2 := individual_rucksack[len(individual_rucksack)/2:]
		for i := 0; i < len(individual_rucksack); i++ {
			c := rune(individual_rucksack[i])
			if i < len(comp1) && strings.ContainsRune(comp2, c) || i >= len(comp1) && strings.ContainsRune(comp2, c) {
				return priorities_maping[c]
			}
		}
		return 0
	}

	//generate index:value pairs for priorities
	priorities := make(map[rune]int)
	for d := 'A'; d <= 'Z'; d++ {
		priorities[d] = int(d - 'A' + 27)
	}
	for c := 'a'; c <= 'z'; c++ {
		priorities[c] = int(c - 'a' + 1)
	}

	//calc total value
	for _, rucksack := range rucksacks {
		total_result += calc_priority(priorities, rucksack)
	}
	fmt.Println(total_result)
	part2day3(rucksacks, priorities)
}

func part2day3(rucksaks []string, mapValue map[rune]int) {
	//anonymous function to calc_priority
	calc_priority := func(priorities_maping map[rune]int, rucksack1 string, rucksack2 string, rucksack3 string) int {
		for i := 0; i < len(rucksack1); i++ {
			c := rune(rucksack1[i])
			if i < len(rucksack1) && strings.ContainsRune(rucksack2, c) && strings.ContainsRune(rucksack3, c) ||
				i >= len(rucksack1) && strings.ContainsRune(rucksack2, c) && strings.ContainsRune(rucksack3, c) {
				return priorities_maping[c]
			}
		}
		return 0
	}

	count := 0
	valueResult := 0

	for {
		valueResult += calc_priority(mapValue, rucksaks[count], rucksaks[count+1], rucksaks[count+2])
		count += 3
		if count >= len(rucksaks) {
			break
		}
	}
	fmt.Println(valueResult)
}

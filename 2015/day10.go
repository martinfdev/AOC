package main

import (
	"fmt"
	"strconv"
)

func Day10() {
	data := Read_file("files/day10.txt")
	// Convert data cons array
	input := []int{}
	for _, v := range data {
		value, _ := strconv.Atoi(string(v))
		input = append(input, value)
	}
	// Part 1
	fmt.Println("Part 1: ", len(part1(input, 40)))
}

// func for const of conway's
func part1(input []int, count_func int) []int {
	if count_func == 0 {
		return input
	}

	output := []int{}
	current := 0
	previos := 0
	count := 0
	for i := 0; i < len(input); i++ {
		current = input[i]
		if current == previos {
			count++
		} else {
			if count > 0 {
				output = append(output, count)
				output = append(output, previos)
			}
			count = 1
			previos = current
		}
	}
	output = append(output, count)
	output = append(output, previos)
	if count_func > 1 {
		return part1(output, count_func-1)
	}
	return part1(output, count_func-1)
}

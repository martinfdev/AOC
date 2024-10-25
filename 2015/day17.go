package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day17() {
	content := Read_file("./files/day17.txt")
	containers := parseContainers(content)

	// Part 1
	const targetVolume = 150
	ways := countWaysDP(containers, targetVolume)
	fmt.Printf("Day 17 - Part 1: Number of ways to fill %d liters: %d\n", targetVolume, ways)
}

func parseContainers(input string) []int {
	var containers []int
	for _, line := range strings.Split(input, "\n") {
		value, _ := strconv.Atoi(line)
		containers = append(containers, value)
	}
	return containers
}

// number of combinations of containers using each container at most once to reach the target volume
func countWaysDP(capacities []int, target int) int {
	ways := make([]int, target+1)
	ways[0] = 1 // one way to reach volume 0

	for _, cap := range capacities {
		for s := target; s >= cap; s-- {
			ways[s] += ways[s-cap]
		}
	}
	return ways[target]
}

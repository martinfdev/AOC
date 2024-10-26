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

	// Part 2
	minK, count := countMinContainerWays(containers, targetVolume)
	fmt.Printf("Day 17 - Part 2: Minimum number of containers: %d, Number of ways using minimum containers: %d\n", minK, count)
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

// how many combinations use the minimum number of containers to reach the target volume
func countMinContainerWays(capacities []int, target int) (minK int, count int) {
	n := len(capacities)
	waysByK := make([][]int, n+1)
	for k := 0; k <= n; k++ {
		waysByK[k] = make([]int, target+1)
	}
	waysByK[0][0] = 1 // one way to reach volume 0 with 0 containers

	for _, cap := range capacities {
		for k := n - 1; k >= 0; k-- {
			for s := target; s >= cap; s-- {
				waysByK[k+1][s] += waysByK[k][s-cap]
			}
		}
	}

	for k := 1; k <= n; k++ {
		if waysByK[k][target] > 0 {
			return k, waysByK[k][target]
		}
	}
	return 0, 0
}

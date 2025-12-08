package main

import (
	"fmt"
	"strings"
)

func Day7() {
	input := read_file("./files/day7.txt")
	parsedInput := parseInput(input)
	fmt.Println("Day 7:", solveTachyon(parsedInput))
}

func parseInput(input string) [][]rune {
	input = strings.TrimSpace(input)
	lines := strings.Split(input, "\n")

	grid := make([][]rune, len(lines))

	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func solveTachyon(grid [][]rune) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	totalSplits := 0

	activeBeams := make(map[int]bool)

	for x := 0; x < cols; x++ {
		if grid[0][x] == 'S' {
			activeBeams[x] = true
			break
		}
	}

	for y := 0; y < rows; y++ {

		nextBeams := make(map[int]bool)

		if len(activeBeams) == 0 {
			break
		}

		for x := range activeBeams {
			currentChar := grid[y][x]

			if currentChar == '^' {
				totalSplits++

				if x-1 >= 0 {
					nextBeams[x-1] = true
				}
				if x+1 < cols {
					nextBeams[x+1] = true
				}
			} else {
				nextBeams[x] = true
			}
		}

		activeBeams = nextBeams
	}
	return totalSplits
}

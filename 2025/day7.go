package main

import (
	"fmt"
	"strings"
)

func Day7() {
	input := read_file("./files/day7.txt")
	parsedInput := parseInput(input)
	fmt.Println("Day 7:", solveTachyon(parsedInput))
	fmt.Println("Day 7.2:", solveQuantumTachyon(parsedInput))
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

func solveQuantumTachyon(grid [][]rune) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])

	activeTimeLines := make(map[int]uint64)

	//find start
	for x := 0; x < cols; x++ {
		if grid[0][x] == 'S' {
			activeTimeLines[x] = 1
			break
		}
	}

	//simulate row by row
	for y := 0; y < rows; y++ {
		nextTimeLines := make(map[int]uint64)

		if len(activeTimeLines) == 0 {
			break
		}

		for x, count := range activeTimeLines {
			currentChar := grid[y][x]

			if currentChar == '^' {
				if x-1 >= 0 {
					nextTimeLines[x-1] += count
				}
				if x+1 < cols {
					nextTimeLines[x+1] += count
				}
			} else {
				nextTimeLines[x] += count
			}
		}

		activeTimeLines = nextTimeLines
	}

	var total uint64 = 0

	for _, count := range activeTimeLines {
		total += count
	}
	return int(total)
}

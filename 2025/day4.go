package main

import (
	"fmt"
	"strings"
)

func Day4() {
	content := read_file("./files/day4.txt")
	result := countAccesibles(content)
	fmt.Println("Day 4 part1: ", result)
	result2 := solveCascading(content)
	fmt.Println("Day 4 part2: ", result2)

}

func countAccesibles(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return 0
	}

	rows := len(lines)
	cols := len(lines[0])

	stride := cols + 2
	grid := make([]bool, (rows+2)*stride)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if lines[r][c] == '@' {
				idx := (r+1)*stride + (c + 1)
				grid[idx] = true
			}
		}
	}

	offsets := [8]int{
		-stride - 1, -stride, -stride + 1, // Fila de arriba
		-1, +1, // Izquierda, Derecha
		stride - 1, stride, stride + 1, // Fila de abajo
	}

	accesibles := 0

	for r := 1; r <= rows; r++ {
		rowStart := r * stride

		for c := 1; c <= cols; c++ {
			currIdx := rowStart + c

			if !grid[currIdx] {
				continue
			}

			neighborCount := 0
			safe := true

			for _, offset := range offsets {
				if grid[currIdx+offset] {
					neighborCount++
					if neighborCount >= 4 {
						safe = false
						break
					}
				}
			}

			if safe {
				accesibles++
			}
		}
	}
	return accesibles
}

func solveCascading(input string) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return 0
	}

	//cascading logic
	rows := len(lines)
	cols := len(lines[0])

	// 1. Configure grid with padding
	stride := cols + 2
	totalSize := (rows + 2) * stride

	grid := make([]bool, totalSize)

	neighborCounts := make([]int8, totalSize)

	//offsets precalculated for the 8 neighbors
	offsets := [8]int{
		-stride - 1, -stride, -stride + 1,
		-1, +1,
		stride - 1, stride, stride + 1,
	}

	// 2. initial loading
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if lines[r][c] == '@' {
				idx := (r+1)*stride + (c + 1)
				grid[idx] = true
			}
		}
	}

	// 3. initial neighbor counting
	queue := make([]int, 0, rows*cols/2)

	for r := 1; r <= rows; r++ {
		rowStart := r * stride
		for c := 1; c <= cols; c++ {
			idx := rowStart + c
			if !grid[idx] {
				continue
			}
			count := int8(0)
			for _, offset := range offsets {
				if grid[idx+offset] {
					count++
				}
			}
			neighborCounts[idx] = count

			if count < 4 {
				queue = append(queue, idx)
			}
		}
	}

	// 4. process in cascade (BFS / event driven)
	totalRemoved := 0

	head := 0 // queue head index

	for head < len(queue) {
		currIdx := queue[head]
		head++

		if !grid[currIdx] {
			continue
		}

		// remove current
		grid[currIdx] = false
		totalRemoved++

		// update neighbors
		for _, offset := range offsets {
			neighborIdx := currIdx + offset

			if grid[neighborIdx] {
				neighborCounts[neighborIdx]--
				if neighborCounts[neighborIdx] < 4 {
					queue = append(queue, neighborIdx)
				}
			}

		}
	}
	return totalRemoved
}

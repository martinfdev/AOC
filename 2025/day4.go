package main

import (
	"fmt"
	"strings"
)

func Day4() {
	content := read_file("./files/day4.txt")
	result := countAccesibles(content)
	fmt.Println("Day 4 part1: ", result)

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

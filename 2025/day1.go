package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day1() {
	content := read_file("./files/day1.txt")
	directions := parse_directions(content)
	fmt.Println(day1_part1(directions, 50))
}

type Direction struct {
	L int
	R int
}

func parse_directions(s string) []Direction {
	direcions := []Direction{}
	parts := strings.Split(s, "\n")
	for _, part := range parts {
		if part[0] == 'L' {
			num, err := strconv.Atoi(part[1:])
			if err != nil {
				continue
			}
			direcions = append(direcions, Direction{L: num, R: -1})

		} else {
			num, err := strconv.Atoi(part[1:])
			if err != nil {
				continue
			}
			direcions = append(direcions, Direction{L: -1, R: num})
		}
	}
	return direcions
}

func day1_part1(directions []Direction, initpoint int) int {
	count := 0
	pos := initpoint
	for _, direction := range directions {
		if direction.L != -1 {
			pos = pos - direction.L
		} else if direction.R != -1 {
			pos = pos + direction.R
		}

		pos = ((pos + 100) + 100) % 100
		if pos == 0 {
			count++
		}
	}
	return count
}

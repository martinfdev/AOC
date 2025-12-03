package main

import (
	"fmt"
	"strings"
)

func Day3() {
	content := read_file("./files/day3.txt")

	total := maxJoltageBank(content)
	fmt.Println("Day 3 Max joltage:", total)
}

func maxJoltageBank(input string) int {
	lines := strings.Split(input, "\n")
	total := 0

	for _, line := range lines {
		n := len(line)
		if n < 2 {
			continue
		}

		bestSecond := int(line[n-1] - '0')
		best := -1
		for i := n - 2; i >= 0; i-- {
			di := int(line[i] - '0')
			cand := 10*di + bestSecond
			if cand > best {
				best = cand
			}
			if di > bestSecond {
				bestSecond = di
			}
		}
		total += best
	}
	return total
}

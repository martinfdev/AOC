package main

import (
	"fmt"
	"strings"
)

func Day3() {
	content := read_file("./files/day3.txt")

	total := maxJoltageBank(content)
	fmt.Println("Day 3 Max joltage:", total)
	k := 12
	total = maxJoltageK(content, k)
	fmt.Println("Day 3 Max joltage with k =", k, ":", total)
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

// part 2 with optimization algorithm and k number
func maxJoltageK(input string, k int) int {
	total := 0

	digitsBuf := make([]byte, 0, 100)
	stack := make([]byte, 0, 100)

	for i := 0; i < len(input); i++ {
		b := input[i]
		if b >= '0' && b <= '9' {
			digitsBuf = append(digitsBuf, b)
		}

		isEnd := (i == len(input)-1)
		if b == '\n' || isEnd {

			if len(digitsBuf) < k {
				digitsBuf = digitsBuf[:0]
				continue
			}

			toDrop := len(digitsBuf) - k
			stack = stack[:0]

			for _, digit := range digitsBuf {

				for toDrop > 0 && len(stack) > 0 && digit > stack[len(stack)-1] {
					stack = stack[:len(stack)-1] // pop
					toDrop--
				}
				stack = append(stack, digit) // push
			}

			stack = stack[:k] // keep only k digits

			lineVal := 0
			for _, dByte := range stack {
				lineVal = lineVal*10 + int(dByte-'0')
			}
			total += lineVal

			digitsBuf = digitsBuf[:0] // reset for next line

		}
	}
	return total
}

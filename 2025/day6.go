package main

import (
	"fmt"
	"strconv"
	"strings"
)

type MathProblem struct {
	Operands []int
	Operator string
}

func Day6() {
	input := read_file("./files/day6.txt")
	result := calculateGrandTotal(parseVerticalFile(input))
	fmt.Println("Day 6 part1:", result)
}

func parseVerticalFile(input string) []MathProblem {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return nil
	}

	//normalize line lengths
	maxWidth := 0
	for _, line := range lines {
		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	grid := make([]string, len(lines))
	for i, line := range lines {
		grid[i] = fmt.Sprintf("%-"+strconv.Itoa(maxWidth)+"s", line)
	}

	var problems []MathProblem
	colStart := 0

	//clean vertical problems
	for col := 0; col <= maxWidth; col++ {
		isSeparator := false
		if col == maxWidth {
			isSeparator = true
		} else {
			allSpaces := true
			for r := 0; r < len(grid); r++ {
				if grid[r][col] != ' ' {
					allSpaces = false
					break
				}
			}
			isSeparator = allSpaces
		}

		if isSeparator {
			if col > colStart {
				prob := extractProblem(grid, colStart, col)
				problems = append(problems, prob)
			}
			colStart = col + 1
		}
	}
	return problems
}

func extractProblem(grid []string, colStart, colEnd int) MathProblem {
	var nums []int
	var op string

	for r := 0; r < len(grid); r++ {
		val := strings.TrimSpace(grid[r][colStart:colEnd])
		if val == "" {
			continue
		}

		//what is this value?
		if val == "+" || val == "-" || val == "*" || val == "/" {
			op = val
		} else {
			num, err := strconv.Atoi(val)
			if err == nil {
				nums = append(nums, num)
			}
		}
	}
	return MathProblem{Operands: nums, Operator: op}
}

func calculateGrandTotal(problems []MathProblem) int {
	total := 0
	for _, prob := range problems {
		if len(prob.Operands) < 2 {
			continue
		}
		result := prob.Operands[0]
		for i := 1; i < len(prob.Operands); i++ {
			switch prob.Operator {
			case "+":
				result += prob.Operands[i]
			case "-":
				result -= prob.Operands[i]

			case "*":
				result *= prob.Operands[i]
			case "/":
				result /= prob.Operands[i]
			}
		}
		total += result
	}
	return total
}

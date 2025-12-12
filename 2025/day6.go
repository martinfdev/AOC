package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type ExtractionStrategy func(grid []string, colStart, colEnd int) MathProblem

type MathProblem struct {
	Operands []int
	Operator string
}

func Day6() {
	input := read_file("./files/day6.txt")
	result := calculateGrandTotal(parseGeneric(input, strategyHorizontal))
	fmt.Println("Day 6 part1:", result)
	result = calculateGrandTotal(parseGeneric(input, StrategyVertical))
	fmt.Println("Day 6 part2:", result)
}

func parseGeneric(input string, extractor ExtractionStrategy) []MathProblem {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return nil
	}

	//normalize line lengths
	maxWidth := 0
	for _, l := range lines {
		if len(l) > maxWidth {
			maxWidth = len(l)
		}
	}

	grid := make([]string, len(lines))
	for i, l := range lines {
		grid[i] = fmt.Sprintf("%-"+strconv.Itoa(maxWidth)+"s", l)
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
				prob := extractor(grid, colStart, col)
				problems = append(problems, prob)
			}
			colStart = col + 1
		}
	}
	return problems
}

// part1
func strategyHorizontal(grid []string, colStart, colEnd int) MathProblem {
	var nums []int
	var op string
	for r := 0; r < len(grid); r++ {
		val := strings.TrimSpace(grid[r][colStart:colEnd])
		if val == "" {
			continue
		}
		if val == "+" || val == "*" || val == "-" || val == "/" {
			op = val
		} else {
			n, _ := strconv.Atoi(val)
			nums = append(nums, n)
		}
	}
	return MathProblem{Operands: nums, Operator: op}
}

// part2
func StrategyVertical(grid []string, start, end int) MathProblem {
	var nums []int
	var op string
	lastRow := len(grid) - 1

	opChunk := grid[lastRow][start:end]
	if strings.Contains(opChunk, "+") {
		op = "+"
	}
	if strings.Contains(opChunk, "*") {
		op = "*"
	}
	if strings.Contains(opChunk, "-") {
		op = "-"
	}
	if strings.Contains(opChunk, "/") {
		op = "/"
	}

	for c := start; c < end; c++ {
		var sb strings.Builder
		for r := 0; r < lastRow; r++ {
			char := rune(grid[r][c])
			if unicode.IsDigit(char) {
				sb.WriteRune(char)
			}
		}
		if sb.Len() > 0 {
			n, _ := strconv.Atoi(sb.String())
			nums = append(nums, n)
		}
	}
	return MathProblem{Operands: nums, Operator: op}
}

func calculateGrandTotal(problems []MathProblem) int {
	total := 0
	for _, p := range problems {
		if len(p.Operands) == 0 {
			continue
		}
		sub := p.Operands[0]
		for i := 1; i < len(p.Operands); i++ {
			if p.Operator == "+" {
				sub += p.Operands[i]
			} else {
				sub *= p.Operands[i]
			}
		}
		total += sub
	}
	return total
}

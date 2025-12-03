package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2() {
	content := read_file("./files/day2.txt")
	ranges := parse_line(content)
	result := sumInvalidIDs(ranges)
	fmt.Printf("Day 2: Sum of invalid IDs: %d\n", result)
}

type Range struct {
	Min int
	Max int
}

func parse_line(content string) []Range {
	var ranges []Range
	arr_ranges := strings.Split(content, ",")
	for _, r := range arr_ranges {
		r = strings.TrimSpace(r)
		if r == "" {
			continue
		}
		bounds := strings.Split(r, "-")
		min, _ := strconv.Atoi(bounds[0])
		max, _ := strconv.Atoi(bounds[1])
		ranges = append(ranges, Range{Min: min, Max: max})
	}
	return ranges
}

func sumInvalidIDs(ranges []Range) int64 {
	var sum int64 = 0

	for _, r := range ranges {
		for id := r.Min; id <= r.Max; id++ {
			if isInvalidID(id) {
				sum += int64(id)
			}
		}
	}
	return sum
}

func isInvalidID(n int) bool {
	if n <= 0 {
		return false
	}
	//count digits has even number of digits
	digits := 0
	tmp := n
	for tmp > 0 {
		digits++
		tmp /= 10
	}

	//if digit is odd, return false
	if digits%2 == 1 {
		return false
	}

	//split number in half
	half := digits / 2
	pow10 := 1
	for i := 0; i < half; i++ {
		pow10 *= 10
	}

	// left = hight part, right = low part
	left := n / pow10
	right := n % pow10

	return left == right
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day4(input_data string) {
	tmp_input := strings.ReplaceAll(input_data, "\r", "")
	pairs := strings.Split(tmp_input, "\n")

	count := 0
	countOverlap := 0

	for _, pair := range pairs {
		asigments := strings.Split(pair, ",")
		first := strings.Split(asigments[0], "-")
		second := strings.Split(asigments[1], "-")

		first_start, _ := strconv.Atoi(first[0])
		first_end, _ := strconv.Atoi(first[1])

		second_start, _ := strconv.Atoi(second[0])
		second_end, _ := strconv.Atoi(second[1])

		if first_start <= second_start && first_end >= second_end ||
			second_start <= first_start && second_end >= first_end {
			count++

			//part2 the solver overlap
		} else if first_start <= second_end && first_end >= second_start {
			countOverlap++
		}
	}
	fmt.Println(count)
	fmt.Println(countOverlap + count)
}

package main

import (
	"regexp"
	"strings"
)

func Day8() {
	input := Read_file("files/day8.txt")
	data_list := strings.Split(input, "\n")
	println("part1: ", count_char(data_list))
}

func count_char(data []string) int {
	literals := 0
	memory := 0

	var double_slash = regexp.MustCompile(`\\\\`)
	var quote = regexp.MustCompile(`\\"`)
	var hex = regexp.MustCompile(`\\x..`)

	for _, line := range data {
		literals += len(line)
		line = double_slash.ReplaceAllString(line, " ")
		line = quote.ReplaceAllString(line, " ")
		line = hex.ReplaceAllString(line, " ")
		memory += len(line) - 2
	}
	return literals - memory
}

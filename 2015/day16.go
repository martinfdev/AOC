package main

import (
	"bufio"
	"fmt"
	"strings"
)

type AuntSue struct {
	ID    int
	Props map[string]int
}

func Day16() {
	content := Read_file("./files/day16.txt")
	aunts := parseAuntSue(content)
	target := targetMFC()

	// Part 1
	sueIDPart1 := findSuePart1(aunts, target)
	fmt.Printf("Day 16 - Part 1: Aunt Sue ID is %d\n", sueIDPart1)
}

func parseAuntSue(content string) []AuntSue {
	var aunts []AuntSue
	scanner := bufio.NewScanner(strings.NewReader(content))

	for scanner.Scan() {
		line := scanner.Text()
		var id int
		props := make(map[string]int)

		// Example line: "Sue 1: goldfish: 6, trees: 9, akitas: 0"
		fmt.Sscanf(line, "Sue %d:", &id)
		parts := strings.Split(line[strings.Index(line, ":")+1:], ",")
		for _, part := range parts {
			var key string
			var value int
			fmt.Sscanf(strings.TrimSpace(part), "%s %d", &key, &value)
			key = strings.TrimSuffix(key, ":")
			props[key] = value
		}

		aunts = append(aunts, AuntSue{ID: id, Props: props})
	}
	return aunts
}

func targetMFC() map[string]int {
	return map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}
}

func findSuePart1(aunts []AuntSue, target map[string]int) int {
	for _, aunt := range aunts {
		match := true
		for key, value := range aunt.Props {
			if target[key] != value {
				match = false
				break
			}
		}
		if match {
			return aunt.ID
		}
	}
	return -1
}

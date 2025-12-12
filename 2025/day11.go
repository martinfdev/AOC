package main

import (
	"fmt"
	"strings"
)

func Day11() {
	input := read_file("./files/day11.txt")
	graph := parseDay11(input)
	fmt.Println("Day 11, Part 1:", countPathsInteractive(graph, "you", "out"))
}

func parseDay11(input string) map[string][]string {
	graph := make(map[string][]string)
	lines := strings.SplitSeq(strings.TrimSpace(input), "\n")
	for line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}
		key := parts[0]
		value := strings.Fields(parts[1])
		graph[key] = value
	}
	return graph
}

func countPathsInteractive(graph map[string][]string, start, end string) int {
	visited := make(map[string]bool)

	var dfs func(node string) int
	dfs = func(node string) int {
		if node == end {
			return 1
		}
		visited[node] = true
		totalPaths := 0
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				totalPaths += dfs(neighbor)
			}
		}
		visited[node] = false
		return totalPaths
	}

	return dfs(start)
}

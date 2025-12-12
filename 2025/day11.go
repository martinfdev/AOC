package main

import (
	"fmt"
	"strings"
)

func Day11() {
	input := read_file("./files/day11.txt")
	graph := parseDay11(input)
	fmt.Println("Day 11, Part 1:", countPathsInteractive(graph, "you", "out"))
	fmt.Println("Day 11, Part 2:", countPathsWithDacFft(graph, "svr", "out"))
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

func countPathsWithDacFft(graph map[string][]string, start, end string) uint64 {
	order, ok := topoSort(graph)
	if !ok {
		// Fallback if the input ever has cycles.
		return countPathsWithDacFftDFS(graph, start, end)
	}

	index := make(map[string]int, len(order))
	for i, node := range order {
		index[node] = i
	}
	startIdx, hasStart := index[start]
	endIdx, hasEnd := index[end]
	if !hasStart || !hasEnd {
		return 0
	}

	dp := make([][4]uint64, len(order))
	startMask := 0
	if start == "dac" {
		startMask |= 1
	}
	if start == "fft" {
		startMask |= 2
	}
	dp[startIdx][startMask] = 1

	for _, node := range order {
		nodeIdx := index[node]
		state := dp[nodeIdx]
		if state[0] == 0 && state[1] == 0 && state[2] == 0 && state[3] == 0 {
			continue
		}
		for _, neighbor := range graph[node] {
			neighborIdx, ok := index[neighbor]
			if !ok {
				continue
			}
			addMask := 0
			if neighbor == "dac" {
				addMask |= 1
			}
			if neighbor == "fft" {
				addMask |= 2
			}
			for mask := 0; mask < 4; mask++ {
				val := state[mask]
				if val == 0 {
					continue
				}
				dp[neighborIdx][mask|addMask] += val
			}
		}
	}

	return dp[endIdx][3]
}

func topoSort(graph map[string][]string) ([]string, bool) {
	indeg := make(map[string]int)
	for node, neighbors := range graph {
		if _, ok := indeg[node]; !ok {
			indeg[node] = 0
		}
		for _, neighbor := range neighbors {
			indeg[neighbor]++
		}
	}

	queue := make([]string, 0, len(indeg))
	for node, degree := range indeg {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	order := make([]string, 0, len(indeg))
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		order = append(order, node)
		for _, neighbor := range graph[node] {
			indeg[neighbor]--
			if indeg[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return order, len(order) == len(indeg)
}

func countPathsWithDacFftDFS(graph map[string][]string, start, end string) uint64 {
	visited := make(map[string]bool)

	var dfs func(node string, mask int) uint64
	dfs = func(node string, mask int) uint64 {
		if node == end {
			if mask == 3 {
				return 1
			}
			return 0
		}
		visited[node] = true
		var total uint64
		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				newMask := mask
				if neighbor == "dac" {
					newMask |= 1
				}
				if neighbor == "fft" {
					newMask |= 2
				}
				total += dfs(neighbor, newMask)
			}
		}
		visited[node] = false
		return total
	}

	startMask := 0
	if start == "dac" {
		startMask |= 1
	}
	if start == "fft" {
		startMask |= 2
	}
	return dfs(start, startMask)
}

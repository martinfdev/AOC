package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Interval struct {
	Start, End int
}

func Day5() {
	data := read_file("./files/day5.txt")
	result := solveIngredients(data)
	fmt.Println("Day 5 part1:", result)
}

func solveIngredients(input string) int {
	parts := strings.Split(input, "\n\n")
	rangeLines := strings.Split(parts[0], "\n")
	queryLines := strings.Split(parts[1], "\n")

	var intervals []Interval

	for _, line := range rangeLines {
		nums := strings.Split(line, "-")
		s, _ := strconv.Atoi(nums[0])
		e, _ := strconv.Atoi(nums[1])
		intervals = append(intervals, Interval{Start: s, End: e})
	}

	// sort
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// merge
	var merged []Interval
	if len(intervals) > 0 {
		current := intervals[0]

		for i := 1; i < len(intervals); i++ {
			next := intervals[i]
			if next.Start <= current.End {
				if next.End > current.End {
					current.End = next.End
				}
			} else {
				merged = append(merged, current)
				current = next
			}
		}
		merged = append(merged, current)
	}

	//proccess queries
	freshCount := 0

	for _, line := range queryLines {
		line := strings.TrimSpace(line)
		if line == "" {
			continue
		}
		id, _ := strconv.Atoi(line)

		if isFresh(id, merged) {
			freshCount++
		}
	}
	return freshCount
}

func isFresh(id int, intervals []Interval) bool {
	idx := sort.Search(len(intervals), func(i int) bool {
		return intervals[i].End >= id
	})

	if idx < len(intervals) && intervals[idx].Start <= id {
		return true
	}

	return false
}

package main

import (
	"fmt"
	"strings"
)

func Day9() {
	input := read_file("./files/day9.txt")
	result := solveLargestRectangle(parseInputTiles(input))
	fmt.Println("Day 9:", result)
}

func parseInputTiles(input string) []Point {
	lines := strings.Split(input, "\n")

	var points []Point
	for _, line := range lines {
		var x, y int
		r := strings.NewReader(line)
		fmt.Fscanf(r, "%d,%d", &x, &y)
		points = append(points, Point{X: x, Y: y})
	}
	return points
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func solveLargestRectangle(points []Point) int {
	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			width := abs(p2.X-p1.X) + 1
			height := abs(p2.Y-p1.Y) + 1

			area := width * height
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

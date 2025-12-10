package main

import (
	"fmt"
	"strings"
)

type Segment struct {
	P1, P2 Point
}

func Day9() {
	input := read_file("./files/day9.txt")
	result := solveLargestRectangle(parseInputTiles(input))
	fmt.Println("Day 9:", result)
	result2 := solveRestrictedRectangle(parseInputTiles(input))
	fmt.Println("Day 9.2:", result2)
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minF(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}
func maxF(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func solveRestrictedRectangle(points []Point) int {
	segments := make([]Segment, len(points))
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		segments[i] = Segment{P1: p1, P2: p2}
	}

	maxArea := 0

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			minX, maxX := min(p1.X, p2.X), max(p1.X, p2.X)
			minY, maxY := min(p1.Y, p2.Y), max(p1.Y, p2.Y)

			width := (maxX - minX) + 1
			height := (maxY - minY) + 1
			area := width * height

			if area <= maxArea {
				continue
			}

			if isValidRectangle(minX, maxX, minY, maxY, segments) {
				maxArea = area
			}
		}
	}

	return maxArea
}

func isValidRectangle(minX, maxX, minY, maxY int, segments []Segment) bool {
	for _, seg := range segments {
		segMinX, segMaxX := min(seg.P1.X, seg.P2.X), max(seg.P1.X, seg.P2.X)
		segMinY, segMaxY := min(seg.P1.Y, seg.P2.Y), max(seg.P1.Y, seg.P2.Y)

		if seg.P1.X == seg.P2.X {
			if segMinX > minX && segMinX < maxX {
				if !(segMaxY <= minY || segMinY >= maxY) {
					return false
				}
			}
		} else {
			if segMinY > minY && segMinY < maxY {
				if !(segMaxX <= minX || segMinX >= maxX) {
					return false
				}
			}
		}
	}

	midX := float64(minX+maxX) / 2.0
	midY := float64(minY+maxY) / 2.0

	intersections := 0
	for _, seg := range segments {
		if seg.P1.X == seg.P2.X {
			y1, y2 := float64(seg.P1.Y), float64(seg.P2.Y)
			x := float64(seg.P1.X)

			if (midY >= minF(y1, y2)) && (midY < maxF(y1, y2)) {
				if x > midX {
					intersections++
				}
			}
		}
	}

	return intersections%2 != 0
}

package main

import (
	"fmt"
	"sort"
	"strings"
)

type Point struct {
	X, Y, Z, ID int
}

type Edge struct {
	U, V, DistSq int
}

type DSU struct {
	parent []int
}

func Day8() {
	input := read_file("./files/day8.txt")
	points := parseInputDSU(input)
	result := solveDSU(points, 1000)
	fmt.Println("Day 8 - part1:", result)
	result2 := solveKruskal(points)
	fmt.Println("Day 8 - part2:", result2)
}

func parseInputDSU(input string) []Point {
	lines := strings.Split(input, "\n")

	var points []Point
	for id, line := range lines {
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		points = append(points, Point{X: x, Y: y, Z: z, ID: id})
	}
	return points
}

func NewDSU(n int) *DSU {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &DSU{parent: p}
}

func (d *DSU) Find(i int) int {
	if d.parent[i] != i {
		d.parent[i] = d.Find(d.parent[i])
	}
	return d.parent[i]
}

func (d *DSU) Union(i, j int) {
	rootI := d.Find(i)
	rootJ := d.Find(j)
	if rootI != rootJ {
		d.parent[rootI] = rootJ
	}
}

func generateAllEdges(points []Point) []Edge {
	var edges []Edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			dx := p1.X - p2.X
			dy := p1.Y - p2.Y
			dz := p1.Z - p2.Z
			distSq := dx*dx + dy*dy + dz*dz

			edges = append(edges, Edge{U: p1.ID, V: p2.ID, DistSq: distSq})
		}
	}
	return edges
}

func calculateProductOfTop3(dsu *DSU, numPoints int) int {
	sizeMap := make(map[int]int)
	for i := 0; i < numPoints; i++ {
		root := dsu.Find(i)
		sizeMap[root]++
	}

	var sizes []int
	for _, size := range sizeMap {
		sizes = append(sizes, size)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	if len(sizes) < 3 {
		return 0
	}
	return sizes[0] * sizes[1] * sizes[2]
}

func solveDSU(points []Point, limitConnections int) int {
	edges := generateAllEdges(points)

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].DistSq < edges[j].DistSq
	})

	dsu := NewDSU(len(points))
	count := 0

	for _, edge := range edges {
		if count >= limitConnections {
			break
		}
		dsu.Union(edge.U, edge.V)
		count++
	}

	return calculateProductOfTop3(dsu, len(points))
}

func solveKruskal(points []Point) int {
	var edges []Edge
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1 := points[i]
			p2 := points[j]

			dx := p1.X - p2.X
			dy := p1.Y - p2.Y
			dz := p1.Z - p2.Z
			distSq := dx*dx + dy*dy + dz*dz

			edges = append(edges, Edge{U: p1.ID, V: p2.ID, DistSq: distSq})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].DistSq < edges[j].DistSq
	})

	dsu := NewDSU(len(points))
	activeComponents := len(points)

	for _, edge := range edges {
		rootU := dsu.Find(edge.U)
		rootV := dsu.Find(edge.V)

		if rootU != rootV {
			dsu.Union(rootU, rootV)
			activeComponents--
			if activeComponents == 1 {
				p1 := points[edge.U]
				p2 := points[edge.V]
				return p1.X * p2.X
			}
		}
	}
	return 0
}

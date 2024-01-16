package main

import (
	"fmt"
	"strings"
)

type Node struct {
	name     string
	distance int
}

type Graph struct {
	nodes map[string]*Node
	arcs  map[Node]map[*Node]int
}

type priorityQueue []*Node

func (pq priorityQueue) Len() int {
	return len(pq)
}

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].distance < pq[j].distance
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Node))
}

func (pq *priorityQueue) Pop() interface{} {
	n := len(*pq)
	node := (*pq)[n-1]
	*pq = (*pq)[0 : n-1]
	return node
}

func generateGraph(nodes *[]Node, data_arr []string) {

}

// dijkstra algorithm
func shortesrPath(nodes []Node) int {
	return 0
}

func Day9() {
	data := Read_file("files/day9.txt")
	data_arr := strings.Split(data, "\n")
	nodes := make([]Node, 0)
	generateGraph(&nodes, data_arr)
	fmt.Println(shortesrPath(nodes))

}

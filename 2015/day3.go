package main

import "fmt"

func Day3() {
	x := 0 // x coordinate
	y := 0 // y coordinate
	data := Read_file("files/day3.txt")
	houses := make(map[string]int)
	houses_gifts := 1
	houses["0,0"] = 1
	for _, c := range data {
		switch c {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		coord := fmt.Sprintf("%d,%d", x, y)
		if _, ok := houses[coord]; !ok {
			houses[coord] = 1
			houses_gifts++
		}
	}
	fmt.Printf("Santa visited %d houses\n", len(houses))

	println(houses_gifts)
}

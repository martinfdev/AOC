package main

import "fmt"

func Day3() {
	x := 0 // x coordinate
	y := 0 // y coordinate
	robot_x := 0
	robot_y := 0
	robo_santa := false
	data := Read_file("files/day3.txt")
	houses := make(map[string]int)
	houses["0,0"] = 1
	for i, c := range data {
		if i%2 == 0 {
			robo_santa = false
		} else {
			robo_santa = true
		}
		switch c {
		case '^':
			if robo_santa {
				robot_y++
			} else {
				y++
			}
		case 'v':
			if robo_santa {
				robot_y--
			} else {
				y--
			}
		case '>':
			if robo_santa {
				robot_x++
			} else {
				x++
			}
		case '<':
			if robo_santa {
				robot_x--
			} else {
				x--
			}
		}
		if robo_santa {
			_, ok := houses[fmt.Sprintf("%d,%d", robot_x, robot_y)]
			if ok {
				houses[fmt.Sprintf("%d,%d", robot_x, robot_y)]++
			} else {
				houses[fmt.Sprintf("%d,%d", robot_x, robot_y)] = 1
			}
		} else {
			_, ok := houses[fmt.Sprintf("%d,%d", x, y)]
			if ok {
				houses[fmt.Sprintf("%d,%d", x, y)]++
			} else {
				houses[fmt.Sprintf("%d,%d", x, y)] = 1
			}
		}
	}
	fmt.Printf("Santa and robot_stanta total visited %d houses\n", len(houses))
}

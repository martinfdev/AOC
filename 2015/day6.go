package main

import (
	"fmt"
	"strings"
)

func Day6() {
	data := Read_file("files/day6.txt")
	list_data := strings.Split(data, "\n")
	grid_lights := make([][]bool, 1000)
	for i := range grid_lights {
		grid_lights[i] = make([]bool, 1000)
	}
	linghts_on := 0
	for _, v := range list_data {
		grid_data := strings.Split(v, " ")
		if grid_data[0] == "turn" {
			if grid_data[1] == "on" {
				x_start, y_start := 0, 0
				x_end, y_end := 0, 0
				fmt.Sscanf(grid_data[2], "%d,%d", &x_start, &y_start)
				fmt.Sscanf(grid_data[4], "%d,%d", &x_end, &y_end)
				for i := x_start; i <= x_end; i++ {
					for j := y_start; j <= y_end; j++ {
						grid_lights[i][j] = true
					}
				}

			} else if grid_data[1] == "off" {
				x_start, y_start := 0, 0
				x_end, y_end := 0, 0
				fmt.Sscanf(grid_data[2], "%d,%d", &x_start, &y_start)
				fmt.Sscanf(grid_data[4], "%d,%d", &x_end, &y_end)
				for i := x_start; i <= x_end; i++ {
					for j := y_start; j <= y_end; j++ {
						grid_lights[i][j] = false
					}
				}
			}
		}
		if grid_data[0] == "toggle" {
			x_start, y_start := 0, 0
			x_end, y_end := 0, 0
			fmt.Sscanf(grid_data[1], "%d,%d", &x_start, &y_start)
			fmt.Sscanf(grid_data[3], "%d,%d", &x_end, &y_end)
			for i := x_start; i <= x_end; i++ {
				for j := y_start; j <= y_end; j++ {
					grid_lights[i][j] = !grid_lights[i][j]
				}
			}
		}
	}
	for _, v := range grid_lights {
		for _, v2 := range v {
			if v2 {
				linghts_on++
			}
		}
	}
	fmt.Println("Day 6:", linghts_on)
}

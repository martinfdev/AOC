package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Day2() {
	data := Read_file("files/day2.txt")
	data_split := strings.Split(data, "\n")

	total_paper := 0
	total_ribbon := 0
	for i, v := range data_split {
		l_w_h := strings.Split(v, "x")
		//cast string to int
		to_int := func(s string) int {
			num, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println(err)
			}
			return num
		}

		dimension := make([]int, len(l_w_h))
		for i, v := range l_w_h {
			dimension[i] = to_int(v)
		}
		if len(l_w_h) == 3 {
			l, w, h := dimension[0], dimension[1], dimension[2]
			total_paper += (2 * l * w) + (2 * w * h) + (2 * h * l) + min(l*w, w*h, h*l)
			total_ribbon += (2 * min(l+w, w+h, h+l)) + (l * w * h)
		} else {
			fmt.Println("Error: ", i, v)
		}
	}
	fmt.Println("Total paper: ", total_paper)
	fmt.Println("Total ribbon: ", total_ribbon)
}

package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func day1(data string) {
	// clean_data := strings.Trim(data, "\r")
	list := strings.Split(data, "\n")
	count := 0
	sum := 0
	numbelf := 0
	total := 0

	slsum := make([]int, 1)

	for _, data := range list {
		data = strings.Trim(data, "\r")
		if data != "" {
			val, err := strconv.Atoi(data)
			if err != nil {
				fmt.Println(err)
			} else {
				sum += val
			}
		} else {
			count++

			if sum > total {
				total = sum
				numbelf = count
			}
			slsum = append(slsum, sum)
			// fmt.Println("Elfo ", count, " ", sum)
			sum = 0
		}
	}
	fmt.Println("Elfo ", numbelf, " ", total)
	part2(slsum)
}

func part2(list []int) {

	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})

	total := 0
	// fmt.Println(list)
	for i := len(list) - 3; i < len(list); i++ {
		total = total + list[i]
	}
	fmt.Println(total)
}

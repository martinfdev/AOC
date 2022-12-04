package main

import (
	"fmt"
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

			if sum >= total {
				total = sum
				numbelf = count
			}
			// fmt.Println("Elfo ", count, " ", sum)
			sum = 0
		}
	}
	fmt.Println("Elfo ", numbelf, " ", total)
}

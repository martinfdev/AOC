package main

import "fmt"

// the problen of parenthesis
func Day1() {
	data := Read_file("files/day1.txt")
	count := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			count++
		} else if data[i] == ')' {
			count--
		}
	}
	fmt.Println(count)
}

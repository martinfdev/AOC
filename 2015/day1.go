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
			if count == -1 {
				fmt.Println(i + 1)
				break
			}
		}
	}
	fmt.Println(count)
}

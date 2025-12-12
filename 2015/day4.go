package main

import (
	"crypto/md5"
	"fmt"
)

func Day4() {
	data := Read_file("files/day4.txt")
	result := md5_hash(data)
	fmt.Println("The lowest positive number that produces a hash with 6 leading zeroes is", result)
}

func md5_hash(data string) int {
	var i int
	for i = 0; ; i++ {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(data+fmt.Sprintf("%d", i))))
		if hash[0:6] == "000000" {
			break
		}
	}
	return i
}

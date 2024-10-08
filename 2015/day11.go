package main

import (
	"fmt"
	"strings"
)

func Day11() {
	data := Read_file("files/day11.txt")
	fmt.Println("Part1", nextPassword(data))
}

func isValid(password string) bool {
	// Rule 2
	if strings.ContainsAny(password, "iol") {
		return false
	}

	// Rule 1
	hasStraight := false
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i]+2 == password[i+2] {
			hasStraight = true
			break
		}
	}

	if !hasStraight {
		return false
	}

	// Rule 3
	pairCount := 0
	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairCount++
			i++
		}
	}
	return pairCount >= 2
}

func incrementPassword(password string) string {
	// converter password to []rune for easy manipulation
	pw := []rune(password)
	for i := len(pw) - 1; i >= 0; i-- {
		if pw[i] == 'z' {
			pw[i] = 'a' // If it's 'z', change it to 'a' and continue with the next character
		} else {
			pw[i]++ // Increment the character
			break
		}
	}
	return string(pw)
}

func nextPassword(password string) string {
	// Next we keep incrementing the password until it's valid
	for {
		password = incrementPassword(password)
		if isValid(password) {
			return password
		}
	}
}

package main

import (
	"fmt"
	"strings"
)

func Day5() {
	data := Read_file("files/day5.txt")
	//rules
	// It contains at least three vowels (aeiou only), like aei, xazegov, or aeiouaeiouaeiou.
	// It contains at least one letter that appears twice in a row, like xx, abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
	// It does not contain the strings ab, cd, pq, or xy, even if they are part of one of the other requirements.
	list_data := strings.Split(data, "\n")
	var nice_strings int
	for _, data_string := range list_data {
		if has_three_vowels(data_string) && has_double_letter(data_string) && !has_bad_strings(data_string) {
			nice_strings++
		}
	}
	fmt.Println("Nice strings:", nice_strings)
}

func has_three_vowels(data string) bool {
	var vowels int
	for _, v := range data {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			vowels++
		}
	}
	if vowels >= 3 {
		return true
	}
	return false
}

func has_double_letter(data string) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i] == data[i+1] {
			return true
		}
	}
	return false
}

func has_bad_strings(data string) bool {
	bad_strings := []string{"ab", "cd", "pq", "xy"}
	for _, v := range bad_strings {
		if strings.Contains(data, v) {
			return true
		}
	}
	return false
}

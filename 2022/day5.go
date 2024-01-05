package main

import (
	"fmt"
	"strings"
)

func day5(input_data string) {
	tmp_data := strings.ReplaceAll(input_data, "\r", "")
	list_dataStack := strings.Split(tmp_data, "\n")

	for _, line := range list_dataStack {
		s := strings.Split(line, " ")
		fmt.Println(len(s), s)
	}
}

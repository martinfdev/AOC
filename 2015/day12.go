package main

import (
	"fmt"
	"encoding/json"
)

func Day12() {
	content := Read_file("files/day12.txt")
	var data interface{}
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	sum := sumNumbers(data)
	fmt.Println("Part1", sum)
	sum2 := sumNumbers2(data)
	fmt.Println("Part2", sum2)
}

//funcion recursiva para sumar todos los numeros en el json
func sumNumbers(data interface{}) int {
	sum := 0
	switch v := data.(type) {
		case float64:
			sum += int(v)
		case []interface{}:
			for _, val := range v {
				sum += sumNumbers(val)
			}
		case map[string]interface{}:
			for _, val := range v {
				sum += sumNumbers(val)
			}
	}
	return sum
}

//part2
func sumNumbers2(data interface{}) int {
	sum := 0
	switch v := data.(type) {
		case float64:
			sum += int(v)
		case []interface{}:
			for _, val := range v {
				sum += sumNumbers2(val)
			}
		case map[string]interface{}:
			for _, item := range v {
				if str, ok := item.(string); ok && str == "red" {
					return 0 // si hay un "red" en el objeto, no sumar nada
				}
			}

			for _, val := range v {
				sum += sumNumbers2(val)
			}
	}
	return sum
}
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

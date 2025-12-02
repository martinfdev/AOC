package main

import (
	"fmt"
	"io"
	"os"
)

func read_file(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buf := make([]byte, 4)
	var data string
	for {
		read_data, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		data += string(buf[:read_data])
	}
	return data
}

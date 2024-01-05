package main

import (
	"io"
	"os"
)

func Read_file(file_name string) string {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := ""
	for {
		var buf [512]byte
		n, err := file.Read(buf[:])
		data += string(buf[:n])
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
	}
	return data
}

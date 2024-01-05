package main

import "os"

func Read_file(file_name string) string {
	file, err := os.Open(file_name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		panic(err)
	}
	return string(data[:count])
}

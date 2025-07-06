package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("read_write_files/example.txt")
	if err != nil {
		fmt.Println("Error opening the file!")
		return
	}
	defer file.Close()

	fmt.Println("File opened successfully!")
}

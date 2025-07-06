package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file.", file)
		return
	}
	defer file.Close()

	// write dat to file
	data := []byte("Hello World!\n")
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file.", err)
	}
	fmt.Println("Data has been written to file successfully!")
}

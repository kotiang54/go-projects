package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("read_write_files/output.txt")
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
		return
	}
	fmt.Println("Data has been written to file successfully!")

	// write to file using WriteString()
	file, err = os.Create("read_write_files/writeString.txt")
	if err != nil {
		fmt.Println("Error creating file.", file)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello Go!\n")
	if err != nil {
		fmt.Println("Error writing to file.", err)
		return
	}
	fmt.Println("Writing to writeString.txt complete!")
}

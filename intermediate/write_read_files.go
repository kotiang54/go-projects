package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Write() and WriteString() functions
	fmt.Println("=== Writing files ===")

	file, err := os.Create("read_write_files/output.txt")
	if err != nil {
		fmt.Println("Error creating file.", file)
		return
	}
	defer file.Close()

	// write data to file
	data := []byte("Hello World!\n\n\nBye World.")
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

	fmt.Println()
	fmt.Println("=== Reading files ===")

	// Read files from the write functions above
	inFile, err := os.Open("read_write_files/output.txt")
	if err != nil {
		fmt.Println("Error opening file!", err)
		return
	}
	defer func() {
		fmt.Println("Closing open file")
		inFile.Close()
	}()

	fmt.Println("File opened successfully!")

	// Read the contents of the opened file
	data = make([]byte, 1024) // buffer to read data into
	_, err = inFile.Read(data)
	if err != nil {
		fmt.Println("Error reading data from file.", err)
		return
	}
	fmt.Println("File content:", string(data))

	// make use of buffio package
	scanner := bufio.NewScanner(inFile)

	// Read line-by-line
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println("Line:", line)
	}
	err = scanner.Err()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
}

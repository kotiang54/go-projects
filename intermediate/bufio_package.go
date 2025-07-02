package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, bufio packages!\nHow are you doing?"))

	// Reading byte slice
	// Read data upto a certain limit
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])

	// ReadString method
	// - reads data upto a certain delimiting character
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading string:", err)
		return
	}
	fmt.Println("Read string:", line)

	// Bufio.Writer
	// for efficient writing of data
	writer := bufio.NewWriter(os.Stdout)

	// writing byte slice
	data = []byte("Hello, bufio package!\n")
	n, err = writer.Write(data)
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", n)

	// Flush the buffer to ensure all data is written to os.Stdout
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
	}

	// Writing string
	str := "This is a string variable"
	n, err = writer.WriteString(str)
	if err != nil {
		fmt.Println("Error writing string:", err)
		return
	}
	fmt.Printf("Wrote %d bytes.\n", n)

	// Flush the buffer
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing writer:", err)
		return
	}
}

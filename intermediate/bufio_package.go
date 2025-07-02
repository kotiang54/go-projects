package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\n"))

	// Reading byte slice
	data := make([]byte, 20)
	n, err := reader.Read(data)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Read %d bytes: %s\n", n, data[:n])
}

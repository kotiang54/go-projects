package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening the file!")
		return
	}
	defer file.Close()
	fmt.Println("File opened successfully!")

	scanner := bufio.NewScanner(file)

	// Keyword to filter lines
	keyword := "important"

	// Read and filter lines
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, keyword) {
			updatedLine := strings.ReplaceAll(line, keyword, "necessary")
			fmt.Println("Filtered line: ", line)
			fmt.Println("Updated line: ", updatedLine)
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error scanning file:", err)
	}

}

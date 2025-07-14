package main

import (
	"fmt"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	tempFile, err := os.CreateTemp("", "temporaryFile")
	checkError(err)

	fmt.Println("Temporary file created:", tempFile.Name())
	defer tempFile.Close()
	// defer os.Remove(tempFile.Name())
}

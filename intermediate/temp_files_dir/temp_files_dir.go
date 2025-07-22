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
	// create temp files using os.CreateTemp()
	tempFile, err := os.CreateTemp("", "temporaryFile")
	checkError(err)

	fmt.Println("Temporary file created:", tempFile.Name())
	defer tempFile.Close()
	// defer os.Remove(tempFile.Name())

	// create temp directory using os.MkdirTemp()
	tempDir, err := os.MkdirTemp("", "GoCourseTempDir")
	checkError(err)

	// defer os.Remove(tempDir)
	fmt.Println("Temporary directory created:", tempDir)
}

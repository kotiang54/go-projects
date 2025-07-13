package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	// relativePath := "./data/file.txt"
	// absolutePath := "/home/user/docs/file.txt"

	// Join paths using filepath.Join()
	joinedPath := filepath.Join("home", "Documents", "downloads", "file.zip")
	fmt.Println("Joined Path:", joinedPath)

	// filepath.Clean()
	normalizedPath := filepath.Clean("./data/../data/file.txt")
	fmt.Println("Normalized Path:", normalizedPath)

	// filepath.Split()
	dir, file := filepath.Split("/home/user/docs/file.txt")
	fmt.Println("File:", file)
	fmt.Println("Dir:", dir)
	fmt.Println(filepath.Base("home/user/docs/"))
}

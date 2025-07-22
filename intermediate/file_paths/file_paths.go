package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	relativePath := "./data/file.txt"
	absolutePath := "/home/user/docs/file.txt"

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

	// Check if existing path is relative or absolute
	fmt.Println("Is relativePath variable absolute? ", filepath.IsAbs(relativePath))
	fmt.Println("Is absolutePath variable absolute? ", filepath.IsAbs(absolutePath))

	// Extract file extensions
	fmt.Println(filepath.Ext(file))
	fmt.Println(strings.TrimSuffix(file, filepath.Ext(file)))

	// Convert absolute path to relative
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/c", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	// Convert relative path to absolute path
	absPath, err := filepath.Abs(relativePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Absolute Path:", absPath)
}

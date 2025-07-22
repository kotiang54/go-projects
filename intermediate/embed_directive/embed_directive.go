package main

// embed directive for embedding static files into Go program"

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed embed_file.txt
var content string

//go:embed embedFolder
var readWriteFolder embed.FS

func main() {
	fmt.Println("Embedded content:", content)

	// embedding a folder
	content, err := readWriteFolder.ReadFile("embedFolder/output.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Embedded file content:", string(content))

	err = fs.WalkDir(readWriteFolder, "embedFolder", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

}

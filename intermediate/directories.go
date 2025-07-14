package intermediate

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// Using os package

	err := os.Mkdir("subdir1", 0755)
	checkError(err)
	// defer os.RemoveAll("subdir1")

	os.WriteFile("subdir1/file", []byte(""), 0755)
	checkError(os.MkdirAll("subdir1/parent/child", 0755))
	checkError(os.MkdirAll("subdir1/parent/child1", 0755))
	checkError(os.MkdirAll("subdir1/parent/child2", 0755))
	checkError(os.MkdirAll("subdir1/parent/child3", 0755))
	os.WriteFile("subdir1/parent/file", []byte(""), 0755)
	os.WriteFile("subdir1/parent/child/file", []byte(""), 0755)

	// ReadDir() function
	// result is a slice fs.DirEntry
	result, err := os.ReadDir("subdir1/parent")
	checkError(err)

	fmt.Println("Reading subdir1/parent")
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	// Change directory
	checkError(os.Chdir("subdir1/parent/child"))
	result, err = os.ReadDir(".")
	checkError(err)

	fmt.Println("")
	fmt.Println("Reading subdir1/parent/child")
	for _, entry := range result {
		fmt.Println(entry.Name(), entry.IsDir(), entry.Type())
	}

	fmt.Println("")
	fmt.Println("Print the current working directory")
	checkError(os.Chdir("../../.."))
	dir, err := os.Getwd()
	checkError(err)
	fmt.Println(dir)

	// Using filepath package
	// filepath.Walk and filepath.WalkDir-(preferred due to performance)
	pathfile := "subdir1/parent/child"
	fmt.Println("")

	err = filepath.WalkDir(pathfile, func(path string, d os.DirEntry, err error) error {
		checkError(err)
		fmt.Println(path)
		return nil
	})
	checkError(err)

	// ros.RemoveAll()
	checkError(os.RemoveAll("subdir1"))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import "os"

func main() {
	err := os.Mkdir("subdir1", 0755)
	checkError(err)

	defer os.RemoveAll("subdir1")
	os.WriteFile("subdir1/file", []byte(""), 0755)
	checkError(os.MkdirAll("subdir1/parent/child", 0755))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

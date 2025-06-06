package main

import "fmt"

func main() {
	message := "Hello, \nGo!"
	rawMessage := `Hello\nGo` // use backticks for raw string literals

	fmt.Println(message)
	fmt.Println(rawMessage)
}

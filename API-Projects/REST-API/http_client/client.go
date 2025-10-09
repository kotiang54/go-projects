package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// Create a new HTTP client
	client := &http.Client{}

	resp, err := client.Get("http://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Print("Error making GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Read and print the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("Error reading response body:", err)
		return
	}
	fmt.Println("Response body:", string(body))
}

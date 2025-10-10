package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// Create a simple HTTP client that makes a GET request to a specified URL
// and prints the response body to the console.

func main() {
	// Create a new HTTP client
	client := &http.Client{Timeout: 10 * time.Second}

	url := "http://jsonplaceholder.typicode.com/posts/1"
	// url := "http://127.0.0.1:3000"
	// url := "https://swapi.dev/api/people/1"

	resp, err := client.Get(url)
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

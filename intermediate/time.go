package main

import (
	"fmt"
	"time"
)

func main() {
	// Current locat time
	fmt.Println(time.Now())

	// Specific time
	specificTime := time.Date(2024, time.June, 30, 12, 0, 0, 0, time.UTC)
	fmt.Println("Specific time:", specificTime)

	// Parse time
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01") //Mon Jan 2 2006 15:04:05 MST 2006
	fmt.Println(parsedTime)
}

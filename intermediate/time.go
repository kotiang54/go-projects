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
	// Go compiler references the time to: Mon Jan 2 2006 15:04:05 MST 2006
	parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")      //Mon Jan 2 2006 15:04:05 MST 2006
	parsedTime1, _ := time.Parse("06-01-02", "20-05-01")         //Mon Jan 2 2006 15:04:05 MST 2006
	parsedTime2, _ := time.Parse("06-1-2 15-04", "20-5-1 18-03") //Mon Jan 2 2006 15:04:05 MST 2006
	fmt.Println(parsedTime)
	fmt.Println(parsedTime1)
	fmt.Println(parsedTime2)

	// Formating time
	t := time.Now()
	fmt.Println(t)
	fmt.Println("Formatted time:", t.Format("Monday 01-02-2006 15-04-05 MST"))

	onedayLater := t.Add(time.Hour * 24)
	fmt.Println("One day later:", onedayLater.Format("01-02-2006 15-04-05 MST"))
	fmt.Println("One day later:", onedayLater.Weekday())

	weekdayLater := time.Now().Add(time.Hour * 24).Weekday() // put all together
	fmt.Println("Weekday later:", weekdayLater)
}

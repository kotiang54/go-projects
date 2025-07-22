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

	loc, _ := time.LoadLocation("Africa/Nairobi")
	t = time.Date(2025, time.June, 27, 8, 05, 00, 00, time.UTC)

	// Convert this to a specific time zone
	tLocal := t.In(loc)

	// Perform rounding
	roundedTime := t.Round(time.Hour)
	// roundedTimeLocal := roundedTime.In(loc)
	roundedTimeLocal := tLocal.Round(time.Hour)

	fmt.Println("Original time (UTC):", t)
	fmt.Println("Original time (Local):", tLocal)
	fmt.Println("Rounded time (UTC):", roundedTime)
	fmt.Println("Rounde time (Local):", roundedTimeLocal)

	// Caluclate durations
	t1 := time.Date(2025, time.June, 27, 8, 50, 00, 00, time.UTC)
	t2 := time.Date(2025, time.June, 27, 17, 50, 00, 00, time.UTC)
	duration := t2.Sub(t1)
	fmt.Println("Duration:", duration)

	// Compare times
	fmt.Println("t2 is after t1?", t2.After(t1))
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// 00:00:00 UTC on Jan 1, 1970

	now := time.Now()
	unixTime := now.Unix()
	fmt.Println("Current Unix Time:", unixTime)

	// convert unix to human readable format
	t := time.Unix(time.Now().Unix(), 0)
	fmt.Println(t)
	fmt.Println("Time:", t.Format("2006-01-02 15-04-05-00 MST"))

	// Mon Jan 02 2006 15:04:05 MST - Go reference date time
	reference := "2006-01-02T15:04:05Z07:00"
	date_str := "2025-06-28T23:40:18Z"

	t, err := time.Parse(reference, date_str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	fmt.Println(t)

	ref := "Jan 02, 2006 03:04 PM"
	str := "Jun 28, 2006 11:50 PM"

	t1, err := time.Parse(ref, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
	}
	fmt.Println(t1)
}

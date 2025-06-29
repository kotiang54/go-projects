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
}

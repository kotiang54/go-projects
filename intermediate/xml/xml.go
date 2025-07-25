package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
	City    string   `xml:"city"`
	Email   string   `xml:"email"`
}

func main() {
	person := Person{
		Name:  "Alice",
		Age:   30,
		City:  "New York",
		Email: "alice.kendagor@example.com",
	}

	xlmData, err := xml.Marshal(person)
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}
	fmt.Println("XML data - person:", string(xlmData))
}

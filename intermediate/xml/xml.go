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
	fmt.Println(string(xlmData))
	fmt.Println("")
	fmt.Println("XML printing with indent:")

	// Using MarshalIndent for pretty printing
	xlmData, err = xml.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}
	fmt.Println(string(xlmData))

	// XML Unmarshalling example
	xmlRawData := `<person><name>Bob</name><age>35</age><city>Chicago</city><email>bob@example.com</email></person>`

	var personData Person
	err = xml.Unmarshal([]byte(xmlRawData), &personData)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return
	}
	fmt.Println("Unmarshalled XML data - person:", personData)
	fmt.Println("Name:", personData.Name)
	fmt.Println("Age:", personData.Age)
	fmt.Println("City:", personData.City)
	fmt.Println("Email:", personData.Email)
	fmt.Println("")
}

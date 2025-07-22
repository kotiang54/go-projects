package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Strings -
	message := "Hello, Go!"
	rawMessage := `Hello\nGo` // use backticks for raw string literals

	fmt.Println(message)
	fmt.Println(rawMessage)

	// Length of strings
	fmt.Println("Length of message variable is", len(message))

	// Compare strings
	str1 := "Apple"
	str2 := "Banana"
	str3 := "apple"
	fmt.Println(str1 < str2) // uses lexicographical comparison - dictionary ordering
	fmt.Println(str1 < str3)

	// String iterations
	for i, char := range message {
		fmt.Printf(" Character at index %d is %c\n", i, char)
	}

	fmt.Println("The first character in message var is", message[0])         // returns ASCII UTF-8 value
	fmt.Println("The first character in message var is", string(message[0])) // returns string character

	fmt.Println("Runes count", utf8.RuneCountInString(message))

	/* Runes - int32 unicode code point: represent individual character in a string
	    - runes are integer value representing a character in Go
		- declared using single quotes
		- handle unicode character useful for multiple languages
	*/
	var ch rune = 'a'
	var jch rune = '日'

	// Print ASCII value
	fmt.Println(ch)
	fmt.Println(jch)

	// Format to the characters
	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	// OR
	fmt.Println(string(ch))
	fmt.Printf("Type of the character %T\n", string(ch))

	jhello := "こんにちは" // Japanese "Hello"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}
}

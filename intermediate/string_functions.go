package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "Hello Go!"
	fmt.Println(len(str))

	str1 := "Hello"
	str2 := "World"
	res := str1 + " " + str2

	fmt.Println(res)
	fmt.Println(str[0])
	fmt.Println(str[1:7])

	// String conversion
	num := 18
	str3 := strconv.Itoa(num)
	fmt.Println(len(str3))

	// String splitting
	fruits := "apple, orange, banana"
	parts := strings.Split(fruits, ",")
	fmt.Println(fruits)
	fmt.Println(parts)

	// Slice of strings - implement join
	countries := []string{"Germany", "France", "Italy"}
	joined := strings.Join(countries, ", ")
	fmt.Println(joined)

	fmt.Println(strings.Contains(str, "Go"))

	// Replace strings
	replaced := strings.Replace(str, "Go", "World", 1)
	fmt.Println(replaced)

	// Trim leading and trailing white spaces
	strwspace := " Hello Everyone! "
	fmt.Println(strwspace)
	fmt.Println(strings.TrimSpace(strwspace))

	// ToUpper() and ToLower()
	fmt.Println(strings.ToLower(strwspace))
	fmt.Println(strings.ToUpper(strwspace))

	// repeat function
	fmt.Println(strings.Repeat("foo ", 3))

	// Count substring, char in a string
	fmt.Println(strings.Count("Hello World!", "l"))

	// Advanced techniques - regular expressions
	str5 := "Hello, 123, Go! 11 743"
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(str5, -1) // -1 return all occurance of matches
	fmt.Println(matches)

	str6 := "Hello, こんにちは"
	fmt.Println(utf8.RuneCountInString(str6))

	// strings.Builder
	var builder strings.Builder

	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("World!")

	// convert builder to  a string
	result := builder.String()
	fmt.Println(result)

	// Using Writerune to add a character
	builder.WriteRune(' ')
	builder.WriteString("How are you!")

	result = builder.String()
	fmt.Println(result)

	// Reset the builder
	builder.Reset()
	builder.WriteString("Starting a fresh")
	result = builder.String()
	fmt.Println(result)
}

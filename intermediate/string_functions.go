package main

import (
	"fmt"
	"strconv"
	"strings"
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
}

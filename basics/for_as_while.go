package basics

import "fmt"

func main() {
	// for as while with break statement
	sum := 0
	for {
		sum += 10
		if sum > 50 {
			break
		}
		fmt.Println("Sum ", sum)
	}

	// for as while with continue statement
	num := 1
	for num <= 10 {
		if num&1 == 0 {
			num++
			continue
		}
		fmt.Println("Odd number: ", num)
		num++
	}
}

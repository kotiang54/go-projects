package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome messages
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen an integer number between 1 and 100!")

	var num int
	attempts := 10
	fmt.Printf("You have %v attempts to guess the correct number.", attempts)

	for attempts > 0 {
		fmt.Printf("\nAttempts left: %d\nEnter your guess: ", attempts)
		_, err := fmt.Scanln(&num)

		if err != nil {
			fmt.Println("Invalid input! Please enter a number!")
			continue
		}

		if num == target {
			fmt.Printf("🎉 Correct! The number was %d.\n", target)
			return
		} else {
			if num < target {
				fmt.Println("Your guess is too low!")
			} else {
				fmt.Println("Your guess is too high!")
			}
		}

		attempts--
	}

	fmt.Printf("\n❌ You've run out of attempts! The number was %d.\n", target)
}

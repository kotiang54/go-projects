package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the Dice Game!")

	for {
		// Show the menu
		fmt.Println("1: Roll the dice")
		fmt.Println("2: Exit")
		fmt.Print("Enter your choice (1 or 2): ")

		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil || (choice != 1 && choice != 2) {
			fmt.Println("Invalid choice, please enter 1 or 2.")
			continue
		}
		if choice == 2 {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}

		die1 := rand.Intn(6) + 1
		die2 := rand.Intn(6) + 1

		// show the results
		fmt.Printf("You rolled a %d and a %d\n", die1, die2)
		fmt.Println("Total:", die1+die2)
		fmt.Println()

		// ask if the user want to roll again
		fmt.Print("Do you want to roll again? (y/n): ")
		var rollAgain string

		_, err = fmt.Scan(&rollAgain)
		if err != nil || (rollAgain != "y" && rollAgain != "n") {
			fmt.Println("Invalid input, assuming no.")
			break
		}
		if rollAgain == "n" {
			fmt.Println("Thanks for playing! Goodbye.")
			break
		}
	}
}

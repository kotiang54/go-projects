package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// os.Args and flag packages

	fmt.Println("Command:", os.Args[0])

	for i, arg := range os.Args {
		fmt.Println("Argument", i, ":", arg)
	}

	// flag
	var (
		name string
		age  int
		male bool
	)

	flag.StringVar(&name, "name", "John", "Name of the user")
	flag.IntVar(&age, "age", 18, "Age of user")
	flag.BoolVar(&male, "male", true, "Gender of the user")

	flag.Parse()

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Male:", male)

	// sub commands
	subcommand1 := flag.NewFlagSet("firstSub", flag.ExitOnError)
	subcommand2 := flag.NewFlagSet("secondSub", flag.ExitOnError)

	firstFlag := subcommand1.Bool("processing", false, "Command processing status")
	secondFlag := subcommand1.Int("bytes", 1024, "Byte length of result")

	flagsc2 := subcommand2.String("language", "Go", "Enter your language")

	if len(os.Args) < 2 {
		fmt.Println("This command requires additional commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "firstSub":
		subcommand1.Parse(os.Args[2:])
		fmt.Println("subCommand1:")
		fmt.Println("processing:", *firstFlag)
		fmt.Println("bytes:", *secondFlag)
	case "secondSub":
		subcommand2.Parse(os.Args[2:])
		fmt.Println("subCommand2:")
		fmt.Println("language", *flagsc2)
	default:
		fmt.Println("Unknown subcommand. Use 'firstSub' or 'secondSub'")
		os.Exit(1)
	}
}

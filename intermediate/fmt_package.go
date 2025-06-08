package intermediate

import "fmt"

func main() {
	// fmt.Print
	fmt.Print("Hello ")
	fmt.Print("World! ")
	fmt.Print("13, 1999")

	// fmt.Println
	fmt.Println("")
	fmt.Println("Hello")
	fmt.Println("World!")
	fmt.Println("13, 1999")

	// fmt.Printf
	name := "John"
	age := 34
	fmt.Println("")
	fmt.Printf("Name: %s, age: %d\n", name, age)
	fmt.Printf("Binary: %b, Hex: %#X\n", age, age)

	// Formating Functions
	// fmt.Sprint
	fmt.Println("")
	s := fmt.Sprint("Hello", "World!", 123, 456)
	fmt.Print(s)

	// fmt.Sprintln
	s = fmt.Sprintln("Hello", "World!", 123, 456)
	fmt.Print(s)
	fmt.Print(s)

	// fmt.Sprintf
	sf := fmt.Sprintf("Name: %s, Age: %d", name, age)
	fmt.Println(sf)

	// Scanning Functions
	var name1 string
	var age1 int

	fmt.Print("Enter your name and age:")
	fmt.Scan(&name1, &age1)

	// Scanln stops scanning at each new line
	// fmt.Scanln(&name1, &age1)

	// fmt.Scanf("%s %d", &name1, &age1)
	fmt.Printf("Name: %s, Age: %d\n", name1, age1)

	// Error Formating Functions
	err := checkAge(15)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}

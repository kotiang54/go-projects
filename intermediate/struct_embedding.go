package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person // embedded struct
	empId  string
	salary float64
}

func main() {
	emp := Employee{
		person: person{name: "John Smith", age: 30},
		empId:  "Emp001",
		salary: 63_000,
	}

	// Access the employee name and age
	// without going through the person struct
	fmt.Println("Name:", emp.name)
	fmt.Println("Age:", emp.age)

	fmt.Println("Employee ID:", emp.empId)
	fmt.Println("Salary ($/yr):", emp.salary)
}

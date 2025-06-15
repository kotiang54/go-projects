package intermediate

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person              // embedded struct: Anonymous field
	employeeInfo person // embedded struct: Named field - requires indirect access trough the person struct
	empId        string
	salary       float64
}

// Method inheritance
func (p person) introduce() {
	// `introduce` is a method of person struct and not of Employee
	fmt.Printf("Hi, I am  %s and I'm %d years old.\n", p.name, p.age)
}

// Overriding a method
func (e Employee) introduce() {
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn $%.2f per year.\n", e.name, e.empId, e.salary)
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

	// introduce is called on Employee instance because person struct
	// is embedded in Employee struct
	emp.introduce()
}

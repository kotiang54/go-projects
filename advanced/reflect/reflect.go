package main

import (
	"fmt"
	"reflect"
)

// Why reflection?
// Reflection in Go is used for inspecting types at runtime,
// enabling dynamic behavior, and working with interfaces.
// It allows developers to write more flexible and reusable code by enabling operations
// on types that are not known until runtime. Common use cases include:
// 		serialization/deserialization,
// 		implementing generic functions, and
// 		building frameworks that require type introspection.

// Structs with reflect package
type Person struct {
	Name string
	Age  int
}

// Working with methods
type Greeter struct{}

func (g Greeter) Greet(name string) string {
	return "Hello, " + name
}

func main() {

	var x float64 = 3.54

	v := reflect.ValueOf(x)
	t := v.Type()

	fmt.Println("Type:", t)
	fmt.Println("Kind:", t.Kind())
	fmt.Println("Is Int:", t.Kind() == reflect.Int)
	fmt.Println("Is Float:", t.Kind() == reflect.Float64)
	fmt.Println("Is String:", t.Kind() == reflect.String)
	fmt.Println("Value:", v.Float())
	fmt.Println("Is zero:", v.IsZero())

	// Another example
	fmt.Println("")
	y := 10
	v1 := reflect.ValueOf(&y).Elem()
	fmt.Println("Value v1:", v1)
	fmt.Println("Original value:", v1.Int())

	v1.SetInt(20)
	fmt.Println("New value:", v1.Int())
	fmt.Println("Original variable:", y)

	var itf interface{} = "Hello"
	v2 := reflect.ValueOf(itf)

	fmt.Println("")
	fmt.Println("Type v2:", v2.Type())
	if v2.Kind() == reflect.String {
		fmt.Println("String value:", v2.String())
	}

	// Struct with reflection
	p := Person{
		Name: "Jane Doe",
		Age:  30,
	}

	v = reflect.ValueOf(p)
	fmt.Println("")

	for i := range v.NumField() {
		fmt.Printf("Field %d: %v\n", i, v.Field(i))
	}

	v1 = reflect.ValueOf(&p).Elem()
	nameField := v1.FieldByName("Name")
	if nameField.CanSet() {
		nameField.SetString("John Doe")
	} else {
		fmt.Println("Cannot set name field")
	}

	// Methods with reflection
	g := Greeter{}
	t = reflect.TypeOf(g)
	v = reflect.ValueOf(g)
	var method reflect.Method

	fmt.Println("")

	fmt.Println("Type:", t)
	for i := range t.NumMethod() {
		method = t.Method(i)
		fmt.Printf("Method %d: %s\n", i, method.Name)
	}

	m := v.MethodByName(method.Name)
	args := []reflect.Value{reflect.ValueOf("Alice")}
	result := m.Call(args)

	fmt.Println("Result:", result[0].String())
}

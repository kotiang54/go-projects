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
	name string
	age  int
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
		name: "Jane Doe",
		age:  30,
	}

	v = reflect.ValueOf(p)
	fmt.Println("")

	for i := range v.NumField() {
		fmt.Printf("Field %d: %v\n", i, v.Field(i))
	}

	v1 = reflect.ValueOf(&p).Elem()
	nameField := v1.FieldByName("name")
	if nameField.CanSet() {
		nameField.SetString("John Doe")
	} else {
		fmt.Println("Cannot set name field")
	}
}

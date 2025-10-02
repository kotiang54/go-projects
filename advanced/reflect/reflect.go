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
}

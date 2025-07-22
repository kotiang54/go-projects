package main

import (
	"fmt"
	"math"
)

//const pi float64 = 3.142

/*
// Interface promotes code reuse, decoupling, and polymorphism
// without relying in inheritance
*/
type geometry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	// Here, rect satisfies the interface because it implements
	// all methods in the interface
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	// Here, circle satisfies the interface because it implements
	// all methods in the interface
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	myPrinter(1, "John", 45.9, true)

	printType(9)
	printType(true)
	printType("John")
	printType(65.7)
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	case bool:
		fmt.Println("Type: Boolean")
	default:
		fmt.Println("Type: Unknown")
	}
}

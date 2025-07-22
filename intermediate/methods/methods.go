package main

import "fmt"

type Rectangle struct {
	length float64
	width  float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

// Method on other types
type MyInt int

func (m MyInt) IsPositive() bool {
	return m > 0
}

// Method with embeddings
type Shape struct {
	Rectangle
}

func main() {
	r := Rectangle{
		length: 10,
		width:  6,
	}
	area := r.Area()
	fmt.Println("Area of the rectangle:", area)

	r.Scale(3)
	area = r.Area()
	fmt.Println("Area of the scaled rectangle:", area)

	var num1 MyInt = -5
	num2 := MyInt(9)

	fmt.Println(num1.IsPositive())
	fmt.Println(num2.IsPositive())

	// embedded methods
	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	fmt.Println(s.Area()) // indirect access
	fmt.Println(s.Rectangle.Area())
}

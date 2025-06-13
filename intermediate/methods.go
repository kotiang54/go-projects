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

func main() {
	r := Rectangle{
		length: 10,
		width:  6,
	}
	area := r.Area()
	fmt.Println("Area of the rectangle:", area)
}

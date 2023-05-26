package main

import (
	"fmt"
	"reflect"
)

type triangle struct {
	height float64
	base   float64
}

// Get area of triangle
func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

type square struct {
	sideLength float64
}

// Get area of square
func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

type shape interface {
	getArea() float64
}

// Print area for a shape
func printArea(s shape) {
	fmt.Print(reflect.TypeOf(s))
	fmt.Println("printArea: ", s.getArea())
}

func main() {
	myTriangle := triangle{height: 10, base: 10}
	mySquare := square{sideLength: 10}

	printArea(myTriangle)
	printArea(mySquare)
}

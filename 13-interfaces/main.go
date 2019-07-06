package main

import (
	"fmt"
	"math"
)

// Shape - Interface Definition
type Shape interface {
	area() float64
	// perimeter() float64
}

// Circle struct
type Circle struct {
	x, y, r float64
}

// Rectangle Struct
type Rectangle struct {
	l, b float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r Rectangle) area() float64 {
	return r.l * r.b
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{0, 0, 5}
	rectangle := Rectangle{10, 15}
	fmt.Printf("CIRCLE Area: %f\n", getArea(circle))
	fmt.Printf("RECTANGLE Area: %f\n", getArea(rectangle))
}

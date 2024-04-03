package main

import (
	"fmt"
	"math"
)

type shapes interface {
	area() float64
	perimeter() float64
}

type square struct {
	lengthofside float64
}

func (sq square) area() float64 {
	return math.Pow(sq.lengthofside, 2) 
}

func (sq square) perimeter() float64 {
	return sq.lengthofside * 4
}

type rectangle struct {
	length float64
	width  float64
}

func (rect rectangle) area() float64 {
	return rect.length * rect.width
}

func (rect rectangle) perimeter() float64 {
	return 2 * (rect.length + rect.width)
}

func main() {
	var s shapes
	s = square{lengthofside: 5}
	fmt.Println("Area of Square: ", s.area())
	fmt.Println("Perimeter of Square: ", s.perimeter())
	s = rectangle{length: 10, width: 5}
	fmt.Println("Area of Rectangle: ", s.area())
	fmt.Println("Perimeter of Rectangle: ", s.perimeter())
}

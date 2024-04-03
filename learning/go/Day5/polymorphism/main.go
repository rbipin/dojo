package main

import "fmt"

type square struct {
	lengthOfASide float32
}

type rectangle struct {
	length float32
	width  float32
}

type circle struct {
	diameter float32
}

type Shapes interface {
	area() float32
	perimeter() float32
}

func (sq square) area() float32 {
	return sq.lengthOfASide * sq.lengthOfASide
}

func (sq square) perimeter() float32 {
	return 4 * sq.lengthOfASide
}

func (rect rectangle) area() float32 {
	return rect.length * rect.width
}

func (rect rectangle) perimeter() float32 {
	return 2 * (rect.length + rect.width)
}

func (cir circle) area() float32 {
	radius := cir.diameter / 2
	return 3.14 * (radius * radius)
}

func (cir circle) perimeter() float32 {
	return cir.diameter/ 2
}

func main() {
	var s Shapes
	s = square {
		lengthOfASide: 2,
	}
	fmt.Println("Square")
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())

	s = rectangle {
		length: 3,
		width: 5,
	}
	fmt.Println("Rectangle")
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
	s = circle {
		diameter: 10,
	}
	fmt.Println("Circle")
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}

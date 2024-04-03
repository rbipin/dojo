package main

import (
	"fmt"
	"math"
)

type square struct {
	lengthofaside float64
	referenceId   int
}

// pointer receiver
func (sq *square) area() float64 {
	fmt.Println("Area Reference Id", sq.referenceId)
	sq.referenceId = 2
	return math.Pow(sq.lengthofaside, 2)
}

// value receiver, this method gets a copy
func (sq square) perimeter() float64 {
	fmt.Println("Perimeter Reference id", sq.referenceId)
	sq.referenceId = 3
	return sq.lengthofaside * 4
}

func main() {
	sq := square{lengthofaside: 5,
		referenceId: 1,
	}
	fmt.Println("Area of the Square: ", sq.area())
	fmt.Println("Perimiter of the Square: ", sq.perimeter()) // Will print reference Id 2

	sq = square{lengthofaside: 6,
		referenceId: 1,
	}
	fmt.Println("Perimiter of the Square: ", sq.perimeter()) // Will print reference Id 1
	fmt.Println("Area of the Square: ", sq.area())           // Will print reference Id 1
}

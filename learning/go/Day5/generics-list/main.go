package main

import (
	"fmt"
)

type Number interface {
	int32 | float32
}

func Sum[T Number] (list []T) T {
	var sum T
	for _, item := range list {
		sum += item
	}
	return sum
}


func main() {
	intItems := []int32 {5, 10, 15, 20}
	sumInt := Sum(intItems)
	fmt.Println("Sum of Int ",sumInt)
	floatItems := []float32 {1.0, 2.2, 1.5, 2.0}
	sumfloat := Sum(floatItems)
	fmt.Println("Sum of Int ",sumfloat)
}
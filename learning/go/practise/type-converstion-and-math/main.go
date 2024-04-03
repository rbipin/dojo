package main

import (
	"fmt"
)

func main() {
	fmt.Println("Add Float & int")
	var firstValue int64 = 5
	var secondValue float64 = 10

	sum := float64(firstValue) + secondValue
	fmt.Println(sum)

	fmt.Println("Adding Floats")
	f1, f2, f3 := 23.5, 10.5, 11.3
	floatSum := f1 + f2 + f3
	fmt.Println(floatSum)
}

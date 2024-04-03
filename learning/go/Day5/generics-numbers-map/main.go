package main

import (
	"fmt"
)

//Method 1
func sumInts [K comparable, V int32] (m map[K]V) V {
	var sum V = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

// Method 2
func sumFloats [K comparable, V float32] (m map[K]V) V {
	var sum V = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

// Method 3 Refactored 1 & 2
func sum [K comparable, V int32 | float32] (m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum;
}

// Further Refactored
type Number interface {
	int32 | float32
}

func sumGenerics [K comparable, V Number] (m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum;
}

func main() {
	intMap := map[string]int32 {
		"A": 1,
		"B": 2,
		"C": 3,
	}
	floatMap := map[string]float32 {
		"A": 1.1,
		"B": 2.345,
		"C": 3.56,
	}
	
	sumInt := sumInts(intMap)
	// Refactored
	sumInt1 := sum(intMap)
	// Further Refactored
	sumInt2 := sumGenerics(intMap)
	fmt.Println("Sum Of Int", sumInt, sumInt1, sumInt2)
	
	sumFloat := sumFloats(floatMap)
	// Refactored
	sumFloat1 := sum(floatMap)
	// Further Refactored
	sumFloat2 := sumGenerics(floatMap)
	fmt.Println("Sum Of Float", sumFloat, sumFloat1, sumFloat2)
}
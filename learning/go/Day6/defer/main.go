package main

import "fmt"

func releaseResource(a *int, b *int) {
	a = nil
	b = nil
	fmt.Printf("Release resource a: %d b: %d", a, b)
}

func add(a int, b int) int {
	return a + b
}

func main() {
	a := 5
	b := 5
	defer releaseResource(&a, &b)
	sum := add(a, b)
	fmt.Println("Sum ", sum)
}
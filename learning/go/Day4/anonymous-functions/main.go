package main

import (
	"fmt"
)

func main() {
	printHelloWorld := func(name string) {
		fmt.Println("Hello World!", name)
	}
	add := func (a int, b int) int {
		return a + b;
	}

	printHelloWorld("Bipin")
	fmt.Println(add(2, 3))
}
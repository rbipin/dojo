package main

import (
	"fmt"
)

func main() {
	count := counter()
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())
	fmt.Println(count())

	newCount := counter()
	fmt.Println(newCount())
	fmt.Println(newCount())
}

func counter() func () int {
	counter:= 0
	return func() int {
		counter++
		return counter
	}
}
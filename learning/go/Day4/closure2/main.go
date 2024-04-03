package main

import (
	"fmt"
)

func main() {
	execute(concatNames, "Bipin", "Radhakrishnan")
}

func concatNames(firstname string, lastname string) {
	fmt.Println(firstname + " " + lastname)
	fmt.Println(executeWithReturn(addNumbers, 3, 5))
}

func addNumbers(a int, b int) int {
	return a + b
}

func execute(f func(string, string), firstname string, lastname string) {
	fmt.Println("Before Execution")
	f(firstname, lastname)
	fmt.Println("After Execution")
}

func executeWithReturn(f func(int, int) int, a int, b int) int {
	fmt.Println("Before Execution")
	result := f(a, b)
	fmt.Println("After Execution")
	return result
}

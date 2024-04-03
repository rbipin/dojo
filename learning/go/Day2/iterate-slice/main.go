package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4}
	slice := []int{1, 2, 3, 4}

	fmt.Println("Printing array...")
	for _, val := range array {
		fmt.Println(val)
	}

	fmt.Println("Printing slice...")
	for _, val := range slice {
		fmt.Println(val)
	}
}

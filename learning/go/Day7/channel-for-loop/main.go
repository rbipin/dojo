package main

import (
	"fmt"
)

func numbers(numberChan chan<- int) {
	for i := 1; i <= 10; i++ {
		numberChan <- i
	}
	close(numberChan)
}

func main() {
	numberChan := make(chan int)
	go numbers(numberChan)
	for n := range numberChan {
		fmt.Println(n)
	}
}

package main

import (
	"fmt"
)

func getNumbers(number chan int) {
	for i := 0; i< 10; i++ {
		number <- i
	}
	close(number)
}

func main() {
	number := make(chan int)
	go getNumbers(number)
	for {
		num, isOpen := <- number
		if (!isOpen) {
			fmt.Println("Channel Closed")
			break;
		}
		fmt.Println(num)
	}
	
}
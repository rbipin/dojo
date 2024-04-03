package main

import (
	"fmt"
	"time"
)

func numbers(nums chan int) {
	for i := 0; i < 5; i++ {
		nums <- i
		fmt.Println("Writing to channel ", i)
	}
	close(nums)
}

func main() {
	nums := make(chan int, 2)
	go numbers(nums)
	time.Sleep(2 * time.Second)
	for n := range nums {
		fmt.Println("Reading ", n)
	}
}

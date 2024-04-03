package main

import (
	"fmt"
)

func main() {
	message := make(chan string, 3)
	message <- "Hey"
	message <- "Hello"
	message <- "there"
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}

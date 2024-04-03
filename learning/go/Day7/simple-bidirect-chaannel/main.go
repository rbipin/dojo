package main

import (
	"fmt"
)

func helloWorld(messageChannel chan string) {
	fmt.Println("Hello World Method!")
	messageChannel <- "Hello World"
}

func main() {
	messageChannel := make(chan string)
	go helloWorld(messageChannel)
	message := <-messageChannel
	fmt.Println("Message: ", message)
}

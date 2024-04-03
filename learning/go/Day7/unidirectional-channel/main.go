package main

import "fmt"

func sayHello(message chan <- string, name string) {
	messagePrefix := "Hello "
	message <- messagePrefix + name
}

func main() {
	message := make(chan string)
	go sayHello(message, "Bipin")
	fmt.Println(<-message)
}
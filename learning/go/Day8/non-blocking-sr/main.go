package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func receive(messages chan string) {
	for {
		msg, open := <-messages
		if !open {
			break
		}
		fmt.Println()
		fmt.Print("Listener: Received ", msg)
	}

}

func receiveNonBlocking(messages chan string) {
	for {
		select {
		case msg, open := <-messages:
			fmt.Println("Listener: Received ", msg)
			if !open {
				break
			}
		default:
			fmt.Println("No Message")
			continue
		}
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimRight(input, "\n")
}

func main() {
	messages := make(chan string)
	go receiveNonBlocking(messages)
	for {
		fmt.Print("Sender: ")
		input := getUserInput()
		if input == "end" {
			fmt.Println("Terminate")
			close(messages)
			break
		}
		select {
		case messages <- input:
			continue
		default:
			continue
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Enter your first value: ")
	firstValue := getUserInput()
	fmt.Printf("Enter your Second value: ")
	secondValue := getUserInput()
	if firstValue == secondValue {
		fmt.Printf("Equal")
		return
	}

	if firstValue > secondValue {
		fmt.Printf("First Value is greater")
		return
	} else {
		fmt.Printf("Second Value is greater")
		return
	}
}

func getUserInput() int64 {
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	value, err := strconv.ParseInt(strings.TrimSpace(inputString), 10, 64)
	if err != nil {
		fmt.Print(err)
		panic("first value is incorrect")
	}
	return value
}

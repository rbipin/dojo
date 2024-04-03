package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
)

func main() {
	input := getUserInput()
	var factorial int64 = 1
	if input == 0 || input == 1 {
		fmt.Println("factorial is: ", factorial)
		return
	} 
	
	for i := input; i>=2; i-- {
		factorial = factorial * i
	}
	fmt.Println("factorial is: ", factorial)
}

func getUserInput() int64 {
	fmt.Println("Enter a number:")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	value, err := strconv.ParseInt(strings.TrimSpace(inputString), 10, 64)
	if err != nil {
		fmt.Println(err)
		panic("Input is invalid")
	}
	return value
}
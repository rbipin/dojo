package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"os"
)

func main() {
	fmt.Printf("Enter value:")
	value := getUserInput()
	isPrime := isPrime(value)
	fmt.Println(isPrime)
}

func isPrime(value int64) bool {
	var count int64
	for count = 2; count <= value - 1; count++ {
		if (value % count) == 0 {
			return false
		}
	}
	return true
}

func getUserInput() int64 {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return value;
}
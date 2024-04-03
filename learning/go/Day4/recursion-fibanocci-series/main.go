package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Number of Fibnocci Number to Generate: ")
	count := getUserInput()
	fmt.Printf("%d %d ", 0, 1)
	FibanocciSeries(0, 1, count)
}

func FibanocciSeries(firstNumber int64, secondNumber int64, count int64) {
	if count == 0 {
		return
	}
	newNumber := firstNumber + secondNumber
	fmt.Printf("%d ", newNumber)
	count--
	FibanocciSeries(secondNumber, newNumber, count)
}

func getUserInput() int64 {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		fmt.Print(err)
		panic(err)
	}
	return value
}

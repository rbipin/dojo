package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Name: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Hello ", input)
	fmt.Print("Enter your Age:")
	ageInput, _ := reader.ReadString('\n')
	age, err := strconv.ParseInt(strings.TrimSpace(ageInput), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Your age is ", age)
}
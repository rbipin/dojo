package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Convert Fahrenheit to Celsius...")
	reader := bufio.NewReader(os.Stdin);
	fmt.Print("Enter Temperature in Fahrenheit:")
	input, _ := reader.ReadString('\n');
	tempFahrenheit, err := strconv.ParseInt(strings.TrimSpace(input), 0, 16)
	if err != nil {
		fmt.Println(err)
	}
	tempCelsius := ((tempFahrenheit -32) * 5/9)
	fmt.Printf("Celsius: %vc", tempCelsius)
}
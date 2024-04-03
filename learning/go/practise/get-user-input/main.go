package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Hello Practise!")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n');
	fmt.Println("Input is: ", input);
}
package main

import "fmt"

func main() {
	printString("Red", "Green", "Blue")
	prefixAndPrint("Color", "Red", "Green", "Blue")
}

func printString(s ...string){
	fmt.Println(s)
}

func prefixAndPrint(prefix string, s ...string) {
	for _, val := range s {
		fmt.Println(prefix, "-", val)
	}
}
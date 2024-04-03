package main

import (
	"fmt"
	"runtime/debug"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
		debug.PrintStack()
	}
}

func createName(firstName *string, lastName *string) string {
	defer recovery()
	if (firstName == nil) {
		panic("First Name is null")
	}
	if (lastName == nil) {
		panic("Last Name is null")
	}
	return *firstName + *lastName
}

func main() {
	firstName := "fName"
	createName(&firstName, nil)
}
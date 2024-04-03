package main

import ()

func createName(firstName *string, lastName *string) string {
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
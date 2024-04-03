package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
}

func newperson(firstname string, lastname string) *person {
	p:= person { firstName: firstname, lastName: lastname}
	return &p
}

func main() {
	np := newperson("John", "Doe")
	fmt.Print(np.firstName)
	fmt.Println("",np.lastName)
	fmt.Println(person {firstName: "Jane", lastName: "Doe"})
}
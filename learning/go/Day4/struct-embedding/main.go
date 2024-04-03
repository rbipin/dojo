package main

import (
	"fmt"
)

type name struct {
	firstName string
	lastName string
}

func (n name) PrintName() string {
	return n.firstName + " " +n.lastName
}

type person struct {
	name
	age uint8
}

func main() {
	n := name {
		firstName: "Bipin",
		lastName: "Radhakrishnan",
	}
	p := person {
		name: n,
		age: 15,
	}
	fmt.Println(p.name.PrintName())
	fmt.Println(p)
}
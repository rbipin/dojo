package main

import (
	"fmt"
	"reflect"
)

type person struct {
	firstName string
	lastName  string
	email     string
}

func main() {
	p := person{
		firstName: "Bipin",
		lastName:  "Radhakrishnan",
		email:     "br@gmail.com",
	}
	showDetails(p)
}

func showDetails(item interface{}) {
	t := reflect.TypeOf(item)
	k := t.Kind()
	v := reflect.ValueOf(item)
	fmt.Println("TypeName: ", t.Name())
	fmt.Println("Kind: ", k)
	fmt.Println("values: ", v)
	fmt.Println("Number of Fields: ", t.NumField())
	fmt.Println("Properties")
	for i := 0; i <= t.NumField()-1; i++ {
		keyValue := fmt.Sprintf("%s : %s", t.Field(i).Name, v.Field(i).String())
		fmt.Println(keyValue)
	}
}

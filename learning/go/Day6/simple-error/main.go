package main

import (
	"fmt"
	"errors"
)

func add(a int, b int) (int, error) {
	return -1, errors.New("Cannot add")
}

func main() {
	a,b := 5, 3
	if _, e := add(a, b); e!= nil {
		fmt.Println(e);
	}
}
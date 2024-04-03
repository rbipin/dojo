package main

import (
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("arg: %d - problem: %s", e.arg, e.prob)
}

func farenheitToCelsius(arg int) (int, error) {
	if (arg < 0) {
		return -1, &argError{arg, "Incorrect arguments"}
	}
	celsius := (arg - 32) * (5/9)
	return celsius, nil
}

func main() {
	farenheit := -1
	if c, e := farenheitToCelsius(farenheit); e != nil {
		fmt.Println("error", e);
	} else {
		fmt.Println("Celsius ", c);
	}
}

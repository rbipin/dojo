// using errors.As and errors.Is
package main

import (
	"errors"
	"fmt"
)


type calculationError struct {
	calculatedValue int
	problem string
}

type argError struct {
	length int
	width int
	problem string
}

func (e calculationError) Error() string {
	return fmt.Sprintf("calculatedvalue: %d problem: %s", e.calculatedValue, e.problem)
}

func (e argError) Error() string {
	return fmt.Sprintf("length: %d, width: %d, problem:%s", e.length, e.width, e.problem)
}

func areaOfRectangle(length int, width int) (int, error) {
	if (length <= 0 || width <= 0) {
		return -1, &argError {0, 0,"Length or Width is less than or equal to zero"}
	}
	
	//dummy error
	if (length > 10) {
		return -1, &calculationError{ -1, "Calculation Error"}
	}
	return length * width, nil;
}

func handleError(e error) {
	if e != nil {
		var argE *argError
		if errors.As(e, &argE) {
			fmt.Println(e)
			return;
		}
		var calcE *calculationError
		if (errors.As(e, &calcE)) {
			fmt.Println(calcE)
		}
	}
}

func main() {
	length, width := 0, 0
	_, e := areaOfRectangle(length, width)
	handleError(e);
	length = 4
	width = 6
	area, e := areaOfRectangle(length, width)
	fmt.Println("Area of Rectagle: ", area)
	length = 11
	width = 25
	_, e = areaOfRectangle(length, width)
	handleError(e)
}
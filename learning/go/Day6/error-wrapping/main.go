package main

import (
	"fmt"
	"errors"
)

var notFound = errors.New("No db records found")
var incorrect = errors.New("Incorrect key")

func getData(key int) error {
	if key == 5 {
		return incorrect
	}
	return notFound;
}

func webService(key int) error {
	if err := getData(key); err != nil {
		return fmt.Errorf("Error %w when calling db", err)
	}
	return nil
}

func main() {
	var keys = [2]int { 1, 5}
	for _, val := range keys {
		if err := webService(val); err != nil {
			if errors.Is(err, notFound) {
				fmt.Println("Not Found", err)
			}
			if errors.Is(err, incorrect) {
				fmt.Println("Incorrect", err)
			}
		}
	}
}
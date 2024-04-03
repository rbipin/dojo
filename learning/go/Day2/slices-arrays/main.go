package main

import (
	"fmt"
	"sort"
)

func main() {
	var array [3]string
	array[0] = "Michigan"
	array[1] = "Texas"
	array[2] = "Ohio"
	fmt.Println(array)

	var newArray = [5]string{"Michigan", "Texas", "Ohio", "New York", "Colorado"}
	fmt.Println(newArray)

	// inferred type
	sliceState := []string{"M", "T", "W"}
	fmt.Println(sliceState)

	var newSlice = make([]string, 3, 5)
	newSlice[0] = "W"
	newSlice[1] = "T"
	newSlice[2] = "M"
	fmt.Println(newSlice)
	newSlice = append(newSlice, "G")
	fmt.Println(newSlice)
	sort.Strings(newSlice)
	fmt.Println(newSlice)
	lastRemoved := append(newSlice[:len(newSlice)-1])
	fmt.Println(lastRemoved)
	firstRemoved := append(newSlice[1:])
	fmt.Println(firstRemoved)
}

package main

import (
	"fmt"
)

func getKeys[K comparable, V any](m map[K]V) []K {
	keys := make([]K,0, len(m))
	for key, _ := range m {
		keys = append(keys, key)
	}
	return keys
}

func main() {
	stateMap := map[string]string{
		"MI": "Michigan",
		"IL": "Illinoi",
		"OH": "Ohio",
	}
	stateMapKeys := getKeys(stateMap)
	fmt.Println(stateMapKeys)

	listOfState := make(map[int]string,3)
	listOfState[0] = "Michigan"
	listOfState[1] = "Illinoi"
	listOfState[2] = "Ohio"

	listOfStateIndex := getKeys(listOfState)
	fmt.Println(listOfStateIndex)
}

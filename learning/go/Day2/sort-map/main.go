package main

import (
	"fmt"
	"sort"
)

func main() {
	statesMap := make(map[string]string)
	statesMap["MI"] = "Michigan"
	statesMap["IL"] = "Illinois"
	statesMap["IN"] = "Indiana"
	statesMap["WI"] = "Wisconsin"
	mapLen := len(statesMap)
	fmt.Println(mapLen)
	keys := make([]string, mapLen)
	i := 0
	for k, v := range statesMap {
		fmt.Println(k, ":", v)
		keys[i] = k
		i++
	}
	fmt.Println("After sorting")
	sort.Strings(keys)
	for k := range keys {
		key := keys[k]
		value := statesMap[key]
		fmt.Println(key, ":", value)
	}
}

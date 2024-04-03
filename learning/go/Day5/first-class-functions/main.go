package main

import (
	"fmt"
)

type order struct {
	orderName string
	total     float32
}

func filter(o []order, f func(o order) bool) []order {
	var filtered []order
	for _, item := range o {
		if f(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func main() {
	// filter function
	filterFunc := func(o order) bool {
		if o.total > 15 {
			return true
		}
		return false
	}
	orders := []order{
		{
			orderName: "Cosmetics",
			total:     55,
		},
		{
			orderName: "GM",
			total:     65,
		},
		{
			orderName: "GM2",
			total:     5,
		},
		{
			orderName: "Milk",
			total:     2.56,
		},
		{
			orderName: "Beverages",
			total:     5.34,
		},
		{
			orderName: "Food",
			total:     35,
		},
	}

	filtered := filter(orders, filterFunc)
	fmt.Println(filtered)
}

package main

import(
	"fmt"
)

func main() {
	i := 10;
	printVal(i);
	printAddress(i);
	printPointerVal(&i);
}

func printVal(val int) {
	fmt.Println("Value: ", val)
}

func printAddress(val int) {
	fmt.Println("Value Address: ", &val)
}

func printPointerVal(val *int) {
	fmt.Println("Print Pointer Val Address: ", val)
	fmt.Println("Print Pointer Val: ", *val)
}

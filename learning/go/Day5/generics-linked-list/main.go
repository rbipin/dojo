package main

import (
	"fmt"
)

type LinkedList[T any] struct {
	node *element[T]
	lastNode *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (list *LinkedList[T]) Push(v T) {
	n := &element[T]{ val: v}
	if list.node == nil {
		list.node = n
		list.lastNode = list.node
	} else {
		list.lastNode.next = n
		list.lastNode = list.lastNode.next
	}
}

func(list * LinkedList[T]) GetAll() []T {
	var elements []T
	node := list.node
	for node != nil {
		elements = append(elements, node.val)
		node = node.next
	}
	return elements
}

func main() {
	linkedList := LinkedList[int] {}
	linkedList.Push(5)
	linkedList.Push(4)
	linkedList.Push(3)
	linkedList.Push(2)
	linkedList.Push(1)

	elements := linkedList.GetAll()
	fmt.Println(elements)
}
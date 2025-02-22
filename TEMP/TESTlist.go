package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {

}

func ListPushFront(l *List, data interface{}) {
	n := &NodeL{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = n, n
		return
	} else {
		n.Next, l.Head = l.Head, n
	}

}

func main() {
	link := &List{}
	ListPushFront(link, "Hello")
	ListPushFront(link, "man")
	ListPushFront(link, "how are you")
	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}

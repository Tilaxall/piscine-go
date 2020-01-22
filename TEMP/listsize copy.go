package main

import "fmt"

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
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

func ListSize(l *List) int {
	result := 0
	for l.Head != nil {
		result++
		l.Head = l.Head.Next
	}
	return result
}

func main() {
	link := &List{}
	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")
	ListPushFront(link, "man")
	ListPushFront(link, "man")
	ListPushFront(link, "man")

	fmt.Println(ListSize(link))
}

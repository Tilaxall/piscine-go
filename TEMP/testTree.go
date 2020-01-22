package main

import "fmt"

type node struct {
	left, right *node
	data        int
}

func insert(root *node, elem int) *node {
	if root == nil {
		return &node{data: elem}
	}
	if elem < root.data {
		root.left = insert(root.left, elem)
	} else {
		root.right = insert(root.right, elem)
	}
	return root
}

func PrintTree(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.data)
	PrintTree(root.left)
	PrintTree(root.right)
}

func main() {
	var root *node
	root = insert(root, 4)
	root = insert(root, 8)
	root = insert(root, 6)
	root = insert(root, 1)
	root = insert(root, 7)
	root = insert(root, 9)
	PrintTree(root)
}

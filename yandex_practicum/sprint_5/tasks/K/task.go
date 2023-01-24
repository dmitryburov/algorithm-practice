package main

import (
	"fmt"
)

// Node
// TODO Comment it before submitting
type Node struct {
	value int
	left  *Node
	right *Node
}

func printRange(root *Node, left int, right int) {
	if root == nil {
		return
	}

	if left <= root.value {
		printRange(root.left, left, right)
	}

	if left <= root.value && right >= root.value {
		fmt.Println(root.value)
	}

	if right >= root.value {
		printRange(root.right, left, right)
	}
}

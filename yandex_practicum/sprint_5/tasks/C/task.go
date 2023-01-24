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

func Solution(root *Node) bool {
	if root.left == nil && root.right == nil {
		return true
	} else if root.left == nil || root.right == nil {
		return false
	}

	return valueLR(root.left) == valueRL(root.right)
}

func valueLR(root *Node) string {
	if root == nil {
		return ""
	}
	return valueLR(root.left) + fmt.Sprint(root.value) + valueLR(root.right)
}

func valueRL(root *Node) string {
	if root == nil {
		return ""
	}
	return valueRL(root.right) + fmt.Sprint(root.value) + valueRL(root.left)
}

package main

// Node
// TODO Comment it before submitting
type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) int {
	if root == nil {
		return 0
	}

	l := Solution(root.left)
	r := Solution(root.right)

	if l > r {
		return l + 1
	}

	return r + 1
}

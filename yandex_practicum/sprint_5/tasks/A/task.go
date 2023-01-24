package main

// Node
// TODO закомментить перед отправкой
type Node struct {
	value int
	left  *Node
	right *Node
}

var maxValue = 0

func Solution(root *Node) int {
	maxValue = 0
	recurse(root)

	return maxValue
}

func recurse(node *Node) {
	if node == nil {
		return
	}

	if node.value > maxValue {
		maxValue = node.value
	}

	if node.left != nil {
		recurse(node.left)
	}
	if node.right != nil {
		recurse(node.right)
	}
}

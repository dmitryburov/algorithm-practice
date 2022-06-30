package main

// Node
// TODO Comment it before submitting
type Node struct {
	value int
	left  *Node
	right *Node
}

func insert(root *Node, key int) *Node {
	if root.value <= key {
		if root.right == nil {
			root.right = &Node{value: key}
		} else {
			insert(root.right, key)
		}
		return root
	} else {
		if root.left == nil {
			root.left = &Node{value: key}
		} else {
			insert(root.left, key)
		}
		return root
	}
}

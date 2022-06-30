package main

// Node
// TODO Comment it before submitting
type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) bool {
	if _, balanced := checkThree(root); !balanced {
		return false
	}

	return true
}

func checkThree(node *Node) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftLen, leftStatus := checkThree(node.left)
	rightLen, rightStatus := checkThree(node.right)
	if !leftStatus || !rightStatus {
		return 0, false
	}

	res := leftLen - rightLen
	if res > 1 || res < -1 {
		return 0, false
	}

	if leftLen >= rightLen {
		return leftLen + 1, true
	}

	return rightLen + 1, true
}

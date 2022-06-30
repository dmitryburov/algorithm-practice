package main

// Node
// TODO закомментить перед отправкой
type Node struct {
	value int
	left  *Node
	right *Node
	size  int
}

func split(node *Node, k int) (*Node, *Node) {
	if node == nil {
		return nil, nil
	}

	ls, rs := TreeSizes(node.left, node.right)

	if ls+1 > k {
		ln, rn := split(node.left, k)
		_, rnSize := TreeSizes(ln, rn)

		node.size = node.size - ls + rnSize
		node.left = rn

		return ln, node
	}

	ln, rn := split(node.right, k-(ls+1))
	lnSize, _ := TreeSizes(ln, rn)

	node.size = node.size - rs + lnSize
	node.right = ln

	return node, rn
}

func TreeSizes(l, r *Node) (ls, rs int) {
	if l != nil {
		ls = l.size
	}
	if r != nil {
		rs = r.size
	}

	return
}

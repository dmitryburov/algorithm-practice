package main

// Node структура узла
// TODO закомментить перед отправкой (если компилятор Make)
type Node struct {
	value int
	left  *Node
	right *Node
}

// remove удаление узла
func remove(node *Node, key int) *Node {
	if node == nil {
		return node
	}

	if node.value == key {
		if node.left == nil && node.right == nil {
			return nil
		}

		if node.left == nil {
			return node.right
		}

		if node.right == nil {
			return node.left
		}

		node.value, node.right = findAndDeleteInTree(node.right)
		return node
	}

	if node.value > key {
		node.left = remove(node.left, key)
	} else {
		node.right = remove(node.right, key)
	}

	return node
}

// findAndDeleteInTree
// если есть и правое и левое поддерево, то ищем в правой самое маленькое значение
// попутно удаляя его и ставя на место правого поддерева новое значение
func findAndDeleteInTree(n *Node) (int, *Node) {
	if n.left == nil {
		if n.right == nil {
			return n.value, nil
		}
		return n.value, n.left
	}

	n.value, n.left = findAndDeleteInTree(n.left)
	return n.value, n
}

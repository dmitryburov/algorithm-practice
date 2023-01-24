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
		return nil
	}

	var currNode = node
	var prevNode *Node

	// присутствует ли ключ в BST
	for currNode != nil && currNode.value != key {
		prevNode = currNode
		if currNode.value < key {
			currNode = currNode.right
		} else {
			currNode = currNode.left
		}
	}
	if currNode == nil {
		return node
	}

	// если есть зотя бы один узел
	if currNode.left == nil || currNode.right == nil {
		var newNode *Node // новый узел

		if currNode.left == nil {
			newNode = currNode.right
		} else {
			newNode = currNode.left
		}

		if prevNode == nil {
			return newNode
		}

		if currNode == prevNode.left {
			prevNode.left = newNode
		} else {
			prevNode.right = newNode
		}

		currNode = nil
	} else {
		var parent, tmp *Node

		tmp = currNode.right
		for tmp.left != nil {
			parent = tmp
			tmp = tmp.left
		}

		if parent != nil {
			parent.left = tmp.right
		} else {
			currNode.right = tmp.right
		}

		currNode.value = tmp.value
		tmp = nil
	}

	return node
}

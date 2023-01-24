package main

// ListNode
// Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

var idxNode = 0

// Solution решение
func Solution(head *ListNode, elem string) int {

	if elem == head.data {
		return idxNode
	}

	if head.next != nil {
		idxNode++
	} else {
		return -1
	}

	return Solution(head.next, elem)
}

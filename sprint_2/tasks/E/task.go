package main

// ListNode
// Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

// Solution решение
func Solution(head *ListNode) *ListNode {
	next := head.next
	head.prev, head.next = head.next, head.prev

	if next != nil {
		return Solution(next)
	}

	return head
}

package main

/*
type ListNode struct {
    data     string
    next  *ListNode
    prev  *ListNode
}*/

func Solution(head *ListNode) *ListNode {
	next := head.next
	head.prev, head.next = head.next, head.prev

	if next != nil {
		return Solution(next)
	}

	return head
}

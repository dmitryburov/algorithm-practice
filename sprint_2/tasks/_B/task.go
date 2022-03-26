package main

import "fmt"

// ListNode
// Comment it before submitting
type ListNode struct {
	data string
	next *ListNode
}

// Solution решение
func Solution(head *ListNode) {
	fmt.Println(head.data)
	if head.next != nil {
		Solution(head.next)
	}
}

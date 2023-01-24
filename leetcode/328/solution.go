package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return head
	}

	var (
		current = head
		prev    = head.Next
		next    = prev
		counter int
	)

	for next.Next != nil {
		counter++

		current.Next = next.Next
		current = next
		next = current.Next
	}

	if counter%2 == 0 {
		current.Next = prev
	} else {
		current.Next = nil
		next.Next = prev
	}

	return head
}

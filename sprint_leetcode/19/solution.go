package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := head, head

	for i := 1; i <= n && right != nil; i++ {
		right = right.Next
	}
	if right == nil {
		return head.Next
	}

	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return head
}

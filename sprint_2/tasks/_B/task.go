package main

import (
	"fmt"
)

func Solution(head *ListNode) {
	fmt.Println(head.data)
	if head.next != nil {
		Solution(head.next)
	}
}

package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItemInput struct {
	head *ListNode
	n    int
}

type TestItem struct {
	input  TestItemInput
	output *ListNode
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := removeNthFromEnd(v.input.head, v.input.n)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", getListView(res), getListView(v.output))
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input: TestItemInput{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val:  5,
								Next: nil,
							},
						},
					},
				},
			}, n: 2},
			output: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val:  5,
							Next: nil,
						},
					},
				},
			},
		},
		{
			input: TestItemInput{head: &ListNode{
				Val:  1,
				Next: nil,
			}, n: 1},
			output: nil,
		},
		{
			input: TestItemInput{head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			}, n: 1},
			output: &ListNode{
				Val: 1,
			},
		},
	}
}

func getListView(node *ListNode) []int {
	var res []int

	for node != nil {
		res = append(res, node.Val)
		node = node.Next
	}

	return res
}

package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItem struct {
	input  *ListNode
	output *ListNode
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := oddEvenList(v.input)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", getListView(res), getListView(v.output))
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			&ListNode{
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
			},
			&ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 5,
							Next: &ListNode{
								Val: 6,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val:  7,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			&ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  4,
										Next: nil,
									},
								},
							},
						},
					},
				},
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

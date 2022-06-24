package main

import (
	"testing"
)

func TestTask(t *testing.T) {
	testTask1(t)
	testTask2(t)
}

func testTask1(t *testing.T) {
	node1 := Node{2, nil, nil}
	node2 := Node{3, &node1, nil}
	node3 := Node{1, nil, &node2}
	node4 := Node{6, nil, nil}
	node5 := Node{8, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node3, &node6}

	newHead := remove(&node7, 10)

	if newHead.value != 5 {
		t.Errorf("\nnewHead.Value: expect 5 got %d", newHead.value)
	}
	if newHead.right != &node5 {
		t.Errorf("\nnewHead.Right: expect %v got %v", &node5, newHead.right)
	}
	if newHead.right.value != 8 {
		t.Errorf("\nnewHead.Right.Value: expect 8 got %d", newHead.right.value)
	}
}

func testTask2(t *testing.T) {
	node1_1 := Node{2, nil, nil}
	node1_2 := Node{3, &node1_1, nil}
	node1_3 := Node{1, nil, &node1_2}
	node1_4 := Node{6, nil, nil}
	node1_n1 := Node{11, nil, nil}
	node1_n2 := Node{16, nil, nil}
	node1_n3 := Node{15, &node1_n1, &node1_n2}
	node1_5 := Node{8, &node1_4, nil}
	node1_6 := Node{10, &node1_5, &node1_n3}
	node1_7 := Node{5, &node1_3, &node1_6}

	newHead := remove(&node1_7, 10)

	if newHead.value != 5 {
		t.Errorf("\nnewHead.Value: expect 5 got %d", newHead.value)
	}
	if newHead.right.value != 11 {
		t.Errorf("\nnewHead.Right.Value: expect 11 got %d", newHead.right.value)
	}
}

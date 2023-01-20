package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItemInput struct {
	nums   []int
	target int
}
type TestItem struct {
	input  TestItemInput
	output int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := search(v.input.nums, v.input.target)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", res, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input:  TestItemInput{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			output: 4,
		},
		{
			input:  TestItemInput{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
			output: -1,
		},
		{
			input:  TestItemInput{nums: []int{1}, target: 0},
			output: -1,
		},
	}
}

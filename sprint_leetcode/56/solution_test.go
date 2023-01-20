package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItemInput struct {
	nums [][]int
}
type TestItem struct {
	input  TestItemInput
	output [][]int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := merge(v.input.nums)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", res, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input:  TestItemInput{nums: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			output: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			input:  TestItemInput{nums: [][]int{{1, 4}, {4, 5}}},
			output: [][]int{{1, 5}},
		},
	}
}

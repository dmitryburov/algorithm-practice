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
	output []int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := twoSum(v.input.nums, v.input.target)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", res, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input:  TestItemInput{nums: []int{2, 7, 11, 15}, target: 9},
			output: []int{0, 1},
		},
		{
			input:  TestItemInput{nums: []int{3, 2, 4}, target: 6},
			output: []int{1, 2},
		},
		{
			input:  TestItemInput{nums: []int{3, 3}, target: 6},
			output: []int{0, 1},
		},
	}
}

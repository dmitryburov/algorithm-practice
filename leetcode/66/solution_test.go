package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItemInput struct {
	nums []int
}
type TestItem struct {
	input  TestItemInput
	output []int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := plusOne(v.input.nums)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", res, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input:  TestItemInput{nums: []int{1, 2, 3}},
			output: []int{1, 2, 4},
		},
		{
			input:  TestItemInput{nums: []int{4, 3, 2, 1}},
			output: []int{4, 3, 2, 2},
		},
		{
			input:  TestItemInput{nums: []int{9}},
			output: []int{1, 0},
		},
	}
}

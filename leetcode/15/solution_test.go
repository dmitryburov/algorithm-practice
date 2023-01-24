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
	output [][]int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := threeSum(v.input.nums)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", res, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input:  TestItemInput{nums: []int{-1, 0, 1, 2, -1, -4}},
			output: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			input:  TestItemInput{nums: []int{0, 1, 1}},
			output: nil,
		},
		{
			input:  TestItemInput{nums: []int{0, 0, 0}},
			output: [][]int{{0, 0, 0}},
		},
	}
}

package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItemInput struct {
	arr1, arr2 []int
}

type TestItem struct {
	input  TestItemInput
	output []int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := mergeSorted(v.input.arr1, v.input.arr2)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", res, v.output)
			}
		})
	}
}
func generateTasks() []TestItem {
	return []TestItem{
		{
			input: TestItemInput{
				arr1: []int{1, 2, 3, 4},
				arr2: []int{5, 6, 7, 8},
			},
			output: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			input: TestItemInput{
				arr1: []int{2, 4, 6},
				arr2: []int{1, 3, 5},
			},
			output: []int{1, 2, 3, 4, 5, 6},
		},
		{
			input: TestItemInput{
				arr1: []int{1, 1, 1},
				arr2: []int{2, 2, 2},
			},
			output: []int{1, 1, 1, 2, 2, 2},
		},
		{
			input: TestItemInput{
				arr1: []int{1},
				arr2: []int{1},
			},
			output: []int{1, 1},
		},
	}
}

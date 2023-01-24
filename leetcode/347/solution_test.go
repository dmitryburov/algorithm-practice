package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItemInput struct {
	arr []int
	k   int
}

type TestItem struct {
	input  TestItemInput
	output []int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := topKFrequent(v.input.arr, v.input.k)
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
				arr: []int{1, 1, 1, 2, 2, 3},
				k:   2,
			},
			output: []int{1, 2},
		},
		{
			input: TestItemInput{
				arr: []int{1, 1, 2, 4, 1, 3, 3, 4, 10},
				k:   3,
			},
			output: []int{1, 4, 3},
		},
	}
}

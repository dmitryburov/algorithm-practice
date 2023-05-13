package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItemInput struct {
	num1, num2 []int
	n1, n2     int
}

type TestItem struct {
	input  TestItemInput
	output []int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			merge(v.input.num1, v.input.n1, v.input.num2, v.input.n2)
			if !reflect.DeepEqual(v.input.num1, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", v.input.num1, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input: TestItemInput{
				num1: []int{1, 2, 3, 0, 0, 0},
				n1:   3,
				num2: []int{2, 5, 6},
				n2:   3,
			},
			output: []int{1, 2, 2, 3, 5, 6},
		},
		{
			input: TestItemInput{
				num1: []int{1},
				n1:   1,
				num2: []int{0},
				n2:   0,
			},
			output: []int{1},
		},
		{
			input: TestItemInput{
				num1: []int{0},
				n1:   0,
				num2: []int{1},
				n2:   1,
			},
			output: []int{1},
		},
	}
}

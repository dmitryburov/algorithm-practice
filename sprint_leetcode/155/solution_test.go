package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItemInput struct {
	command string
	value   int
}

type TestItem struct {
	input  []TestItemInput
	output []int
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := realiseTask(v.input)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", res, v.output)
			}
		})
	}
}

func realiseTask(commands []TestItemInput) []int {
	var res []int
	var stack = Constructor()

	for _, v := range commands {
		switch v.command {
		case "push":
			stack.Push(v.value)
		case "pop":
			stack.Pop()
		case "top":
			res = append(res, stack.Top())
		case "min":
			res = append(res, stack.GetMin())
		default:
			continue
		}
	}

	return res
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input: []TestItemInput{
				{
					command: "push",
					value:   -2,
				},
				{
					command: "push",
					value:   0,
				},
				{
					command: "push",
					value:   -3,
				},
				{
					command: "min",
				},
				{
					command: "pop",
				},
				{
					command: "top",
				},
				{
					command: "min",
				},
			},
			output: []int{-3, 0, -2},
		},
	}
}

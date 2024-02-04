package main

import (
	"fmt"
	"testing"
)

type TestItemInput struct {
	t1,
	t2 int
	action string
}

type TestItemOutput int

type TestItem struct {
	input  TestItemInput
	output TestItemOutput
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			val := solution(v.input.t1, v.input.t2, v.input.action)
			if val != int(v.output) {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%d \nВерно: \n%d", val, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{TestItemInput{10, 20, "heat"}, 20},
		{TestItemInput{10, 20, "freeze"}, 10},
		{TestItemInput{10, 20, "auto"}, 20},
		{TestItemInput{-3, -11, "freeze"}, -11},
		{TestItemInput{1, 2, "fan"}, 1},
		{TestItemInput{0, 10, "heat"}, 10},
		{TestItemInput{20, 15, "heat"}, 20},
	}
}

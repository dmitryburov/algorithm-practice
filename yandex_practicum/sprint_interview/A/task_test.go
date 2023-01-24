package main

import (
	"fmt"
	"testing"
)

type TestItemInput struct {
	a, b int
}
type TestItem struct {
	input  struct{ a, b int }
	output int
}

func TestTask(t *testing.T) {
	tasks := generateTasks()

	for i, v := range tasks {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := solution(v.input.a, v.input.b)
			if res != v.output {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%v \nВерно: \n%v", res, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input:  TestItemInput{a: 12, b: 90},
			output: 102,
		},
		{
			input:  TestItemInput{a: 200, b: -200},
			output: 0,
		},
		{
			input:  TestItemInput{a: 1000000000, b: 1000000000},
			output: 2000000000,
		},
	}
}

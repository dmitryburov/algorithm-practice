package main

import (
	"fmt"
	"testing"
)

type TestItemInput struct {
	a,
	b,
	c int
}

type TestItemOutput string

type TestItem struct {
	input  TestItemInput
	output TestItemOutput
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			val := solution(v.input.a, v.input.b, v.input.c)
			if val != string(v.output) {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%s \nВерно: \n%s", val, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{TestItemInput{3, 4, 5}, "YES"},
		{TestItemInput{3, 5, 4}, "YES"},
		{TestItemInput{4, 5, 3}, "YES"},
		{TestItemInput{1, 2, 3}, "NO"},
	}
}

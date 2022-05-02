package main

import (
	"fmt"
	"strings"
	"testing"
)

type task struct {
	N      int
	Input  []string
	Result []string
}

func TestTask(t *testing.T) {
	tasks := generateTasks()
	for i := 0; i < len(tasks); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := testMain(tasks[i].N, tasks[i].Input)
			if strings.Join(res, "") != strings.Join(tasks[i].Result, "") {
				t.Errorf("Неверный ответ решения!\nОтвет: %v \nВерно: %v", res, tasks[i].Result)
			}
		})
	}
}

func generateTasks() []*task {
	return []*task{
		{
			N: 10,
			Input: []string{
				"get 1",
				"put 1 10",
				"put 2 4",
				"get 1",
				"get 2",
				"delete 2",
				"get 2",
				"put 1 5",
				"get 1",
				"delete 2",
			},
			Result: []string{
				"None",
				"10",
				"4",
				"4",
				"None",
				"5",
				"None",
			},
		},
		{
			N: 8,
			Input: []string{
				"get 9",
				"delete 9",
				"put 9 1",
				"get 9",
				"put 9 2",
				"get 9",
				"put 9 3",
				"get 9",
			},
			Result: []string{
				"None",
				"None",
				"1",
				"2",
				"3",
			},
		},
	}
}

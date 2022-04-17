package main

import (
	"fmt"
	"testing"
)

type task struct {
	K      int
	Input  []int
	Result int
}

func TestTask(t *testing.T) {
	tasks := generateTasks()
	for i := 0; i < len(tasks); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := brokenSearch(tasks[i].Input, tasks[i].K)
			if res != tasks[i].Result {
				t.Errorf("Неверный ответ решения!\nОтвет: %d \nВерно: %d", res, tasks[i].Result)
			}
		})
	}
}

func generateTasks() []task {
	return []task{
		{
			K:      5,
			Input:  []int{19, 21, 100, 101, 1, 4, 5, 7, 12},
			Result: 6,
		},
		{
			K:      1,
			Input:  []int{5, 1},
			Result: 1,
		},
		{
			K:      5,
			Input:  []int{0, 2, 6, 7, 8, 9, 10},
			Result: -1,
		},
	}
}

package main

import (
	"fmt"
	"testing"
)

type task struct {
	N, X   int
	Days   []int
	Result []int
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := findDays(taskItems[i].N, taskItems[i].X, taskItems[i].Days)
			if fmt.Sprint(res) != fmt.Sprint(taskItems[i].Result) {
				t.Errorf("Неверный ответ решения!\nОтвет: %v \nВерно: %v", res, taskItems[i].Result)
			} else {
				t.Log("OK")
			}
		})
	}
}

// generateTasks создает задачи для теста
func generateTasks() (tasks []task) {
	tasks = append(
		tasks,
		task{
			N:      6,
			X:      10,
			Days:   []int{1, 2, 4, 4, 4, 4},
			Result: []int{-1, -1},
		},
		task{
			N:      7,
			X:      3,
			Days:   []int{1, 1, 4, 4, 4, 4, 8},
			Result: []int{3, 7},
		},
		task{
			N:      6,
			X:      1,
			Days:   []int{1, 2, 4, 4, 4, 4},
			Result: []int{1, 2},
		},
		task{
			N:      17,
			X:      28,
			Days:   []int{12, 19, 29, 31, 33, 37, 52, 62, 66, 68, 71, 75, 75, 88, 90, 93, 98},
			Result: []int{3, 8},
		},
	)

	return
}

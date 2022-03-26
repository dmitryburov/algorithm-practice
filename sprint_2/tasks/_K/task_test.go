package main

import (
	"fmt"
	"testing"
)

type task struct {
	Input, Result int
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := solution(taskItems[i].Input)
			if res != taskItems[i].Result {
				t.Errorf("Неверный ответ решения!\nОтвет: %d \nВерно: %d", res, taskItems[i].Result)
			}
		})
	}
}

// generateTasks создает задачи для теста
func generateTasks() (tasks []task) {
	tasks = append(
		tasks,
		task{
			Input:  3,
			Result: 3,
		},
		task{
			Input:  0,
			Result: 1,
		},
	)

	return
}

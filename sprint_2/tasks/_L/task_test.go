package main

import (
	"fmt"
	"testing"
)

type task struct {
	Input  []int
	Result int
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := solution(taskItems[i].Input[0], taskItems[i].Input[1])
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
			Input:  []int{3, 1},
			Result: 3,
		},
		task{
			Input:  []int{10, 1},
			Result: 9,
		},
		task{
			Input:  []int{237, 7},
			Result: 471519,
		},
		task{
			Input:  []int{1000000, 8},
			Result: 26937501,
		},
	)

	return
}

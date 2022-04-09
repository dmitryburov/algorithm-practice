package main

import (
	"fmt"
	"testing"
)

type task struct {
	N      int
	Result []string
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := solution(taskItems[i].N, "", 0, 0)
			if fmt.Sprint(res) != fmt.Sprint(taskItems[i].Result) {
				t.Errorf("Неверный ответ решения!\nОтвет: %s \nВерно: %s", res, taskItems[i].Result)
			}
		})
	}
}

// generateTasks создает задачи для теста
func generateTasks() (tasks []task) {
	tasks = append(
		tasks,
		task{
			N: 3,
			Result: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		task{
			N: 2,
			Result: []string{
				"(())",
				"()()",
			},
		},
	)

	return
}

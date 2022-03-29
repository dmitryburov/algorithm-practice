package main

import (
	"fmt"
	"strings"
	"testing"
)

type task struct {
	N      int
	Input  []command
	Result []string
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := solution(taskItems[i].N, taskItems[i].Input)
			if strings.Join(res, "") != strings.Join(taskItems[i].Result, "") {
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
			N: 10,
			Input: []command{
				{
					action: "put",
					num:    -34,
				},
				{
					action: "put",
					num:    -23,
				},
				{
					action: "get",
				},
				{
					action: "size",
				},
				{
					action: "get",
				},
				{
					action: "size",
				},
				{
					action: "get",
				},
				{
					action: "get",
				},
				{
					action: "put",
					num:    80,
				},
				{
					action: "size",
				},
			},
			Result: []string{"-34", "1", "-23", "0", "error", "error", "1"},
		},
		task{
			N: 6,
			Input: []command{
				{
					action: "put",
					num:    -66,
				},
				{
					action: "put",
					num:    98,
				},
				{
					action: "size",
				},
				{
					action: "size",
				},
				{
					action: "get",
				},
				{
					action: "get",
				},
			},
			Result: []string{"2", "2", "-66", "98"},
		},
	)

	return
}

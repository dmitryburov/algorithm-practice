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
					Name: "pop",
				},
				{
					Name: "pop",
				},
				{
					Name: "push",
					Step: 4,
				},
				{
					Name: "push",
					Step: -5,
				},
				{
					Name: "push",
					Step: 7,
				},
				{
					Name: "pop",
				},
				{
					Name: "pop",
				},
				{
					Name: "get_max",
				},
				{
					Name: "pop",
				},
				{
					Name: "get_max",
				},
			},
			Result: []string{"error", "error", "4", "None"},
		},
		task{
			N: 12,
			Input: []command{
				{
					Name: "pop",
				},
				{
					Name: "get_max",
				},
				{
					Name: "pop",
				},
				{
					Name: "pop",
				},
				{
					Name: "pop",
				},
				{
					Name: "push",
					Step: -6,
				},
				{
					Name: "pop",
				},
				{
					Name: "get_max",
				},
				{
					Name: "pop",
				},
				{
					Name: "pop",
				},
				{
					Name: "push",
					Step: -6,
				},
				{
					Name: "get_max",
				},
			},
			Result: []string{"error", "None", "error", "error", "error", "None", "error", "error", "-6"},
		},
	)

	return
}

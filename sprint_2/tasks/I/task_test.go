package main

import (
	"fmt"
	"strings"
	"testing"
)

type task struct {
	N      int
	Input  *queue
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
			N: 8,
			Input: &queue{
				stack: make([]int, 0, 2),
				commands: []command{
					{
						Name: "peek",
					},
					{
						Name: "push",
						Step: 5,
					},
					{
						Name: "push",
						Step: 2,
					},
					{
						Name: "peek",
					},
					{
						Name: "size",
					},
					{
						Name: "size",
					},
					{
						Name: "push",
						Step: 1,
					},
					{
						Name: "size",
					},
				},
				max: 2,
			},
			Result: []string{"None", "5", "2", "2", "error", "2"},
		},
		task{
			N: 10,
			Input: &queue{
				stack: make([]int, 0, 1),
				commands: []command{
					{
						Name: "push",
						Step: 1,
					},
					{
						Name: "size",
					},
					{
						Name: "push",
						Step: 3,
					},
					{
						Name: "size",
					},
					{
						Name: "push",
						Step: 1,
					},
					{
						Name: "pop",
					},
					{
						Name: "push",
						Step: 1,
					},
					{
						Name: "pop",
					},
					{
						Name: "push",
						Step: 3,
					},
					{
						Name: "push",
						Step: 3,
					},
				},
				max: 1,
			},
			Result: []string{"1", "error", "1", "error", "1", "1", "error"},
		},
	)

	return
}

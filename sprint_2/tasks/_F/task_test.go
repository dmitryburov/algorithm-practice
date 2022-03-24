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
		//task{
		//	N:      8,
		//	Input:  []command{
		//		{
		//			Name: "get_max",
		//		},
		//		{
		//			Name: "push",
		//			Step: 7,
		//		},
		//		{
		//			Name: "pop",
		//		},
		//		{
		//			Name: "push",
		//			Step: -2,
		//		},
		//		{
		//			Name: "push",
		//			Step: -1,
		//		},
		//		{
		//			Name: "pop",
		//		},
		//		{
		//			Name: "get_max",
		//		},
		//		{
		//			Name: "get_max",
		//		},
		//	},
		//	Result: []string{"None", "-2", "-2"},
		//},
		//task{
		//	N:      7,
		//	Input:  []command{
		//		{
		//			Name: "get_max",
		//		},
		//		{
		//			Name: "pop",
		//		},
		//		{
		//			Name: "pop",
		//		},
		//		{
		//			Name: "pop",
		//		},
		//		{
		//			Name: "push",
		//			Step: 10,
		//		},
		//		{
		//			Name: "get_max",
		//		},
		//		{
		//			Name: "push",
		//			Step: -9,
		//		},
		//	},
		//	Result: []string{"None", "error", "error", "error", "10"},
		//},
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
					Step: 9,
				},
				{
					Name: "push",
					Step: -7,
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
					Step: -8,
				},
				{
					Name: "get_max",
				},
				{
					Name: "get_max",
				},
			},
			Result: []string{"error", "error", "9", "-8", "-8"},
		},
	)

	return
}

func getNumOperation(n int) *int {
	return &n
}

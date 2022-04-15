package main

import (
	"fmt"
	"strings"
	"testing"
)

type task struct {
	Input  string
	Result string
}

func TestTask(t *testing.T) {
	var taskItems = generateTasks()
	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := solution(taskItems[i].Input)
			if strings.Trim(fmt.Sprint(res), "[]") != taskItems[i].Result {
				t.Errorf("Неверный ответ решения!\nОтвет: %s \nВерно: %s", res, taskItems[i].Result)
			}
		})
	}
}

func generateTasks() (tasks []task) {
	tasks = append(
		tasks,
		task{
			Input:  "23",
			Result: "ad ae af bd be bf cd ce cf",
		},
		task{
			Input:  "92",
			Result: "wa wb wc xa xb xc ya yb yc za zb zc",
		},
	)

	return
}

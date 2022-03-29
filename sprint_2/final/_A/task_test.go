package main

import (
	"fmt"
	"strings"
	"testing"
)

type task struct {
	Deque    *Deque
	Commands []command
	Result   []string
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := taskItems[i].Deque.Run(taskItems[i].Commands)
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
			Deque: &Deque{
				stack:   make([]int, 10),
				maxSize: 10,
			},
			Commands: []command{
				{
					action: "push_front",
					num:    -855,
				},
				{
					action: "push_front",
					num:    0,
				},
				{
					action: "pop_back",
				},
				{
					action: "pop_back",
				},
				{
					action: "push_back",
					num:    844,
				},
				{
					action: "pop_back",
				},
				{
					action: "push_back",
					num:    823,
				},
			},
			Result: []string{"-855", "0", "844"},
		},
		task{
			Deque: &Deque{
				stack:   make([]int, 6),
				maxSize: 6,
			},
			Commands: []command{
				{
					action: "push_front",
					num:    -201,
				},
				{
					action: "push_back",
					num:    959,
				},
				{
					action: "push_back",
					num:    102,
				},
				{
					action: "push_front",
					num:    20,
				},
				{
					action: "pop_front",
				},
				{
					action: "pop_back",
				},
			},
			Result: []string{"20", "102"},
		},
		task{
			Deque: &Deque{
				stack:   make([]int, 7),
				maxSize: 7,
			},
			Commands: []command{
				{
					action: "pop_front",
				},
				{
					action: "pop_front",
				},
				{
					action: "push_front",
					num:    741,
				},
				{
					action: "push_front",
					num:    648,
				},
				{
					action: "pop_front",
				},
				{
					action: "pop_back",
				},
				{
					action: "pop_front",
				},
			},
			Result: []string{"error", "error", "648", "741", "error"},
		},
	)

	return
}

package main

import (
	"fmt"
	"testing"
)

type task struct {
	N, M          int
	Input, Result [][]string
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := solution(taskItems[i].N, taskItems[i].M, taskItems[i].Input, true)
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
			N: 4,
			M: 3,
			Input: [][]string{
				[]string{"1", "2", "3"},
				[]string{"0", "2", "6"},
				[]string{"7", "4", "1"},
				[]string{"2", "7", "0"},
			},
			Result: [][]string{
				[]string{"1", "0", "7", "2"},
				[]string{"2", "2", "4", "7"},
				[]string{"3", "6", "1", "0"},
			},
		},
		task{
			N: 9,
			M: 5,
			Input: [][]string{
				[]string{"-7", "-1", "0", "-4", "-9"},
				[]string{"5", "-1", "2", "2", "9"},
				[]string{"3", "1", "-8", "-1", "-7"},
				[]string{"9", "0", "8", "-8", "-1"},
				[]string{"2", "4", "5", "2", "8"},
				[]string{"-7", "10", "0", "-4", "-8"},
				[]string{"-3", "10", "-7", "10", "3"},
				[]string{"1", "6", "-7", "-5", "9"},
				[]string{"-1", "9", "9", "1", "9"},
			},
			Result: [][]string{
				[]string{"-7", "5", "3", "9", "2", "-7", "-3", "1", "-1"},
				[]string{"-1", "-1", "1", "0", "4", "10", "10", "6", "9"},
				[]string{"0", "2", "-8", "8", "5", "0", "-7", "-7", "9"},
				[]string{"-4", "2", "-1", "-8", "2", "-4", "10", "-5", "1"},
				[]string{"-9", "9", "-7", "-1", "8", "-8", "3", "9", "9"},
			},
		},
	)

	return
}

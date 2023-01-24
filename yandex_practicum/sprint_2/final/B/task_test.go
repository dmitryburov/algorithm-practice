package main

import (
	"fmt"
	"testing"
)

type task struct {
	Input,
	Result string
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res, err := solution(taskItems[i].Input)
			if err != nil {
				t.Errorf("Ошибка выполнения: %s ", err.Error())
			} else if fmt.Sprint(res) != taskItems[i].Result {
				t.Errorf("Неверный ответ решения!\nОтвет: %f \nВерно: %s", res, taskItems[i].Result)
			}
		})
	}
}

// generateTasks создает задачи для теста
func generateTasks() (tasks []task) {
	tasks = append(
		tasks,
		task{
			Input:  "2 1 + 3 *",
			Result: "9",
		},
		task{
			Input:  "7 2 + 4 * 2 +",
			Result: "38",
		},
		task{
			Input:  "4 2 * 4 / 25 * 2 - 12 / 1000 + 2 / -999 +",
			Result: "-497",
		},
		task{
			Input:  "4 2 * 4 / 25 * 2 - 12 / 500 2 * + 2 / -999 + 71 + -1 * 2 / 1000 + 6 * 8065 -",
			Result: "-787",
		},
		task{
			Input:  "4 2 * 4 / 25 * 2 - 12 / 500 2 * + 2 / -999 + 71 + -1 * 2 / 1000 + 6 * 8065 - 787 + 66 *",
			Result: "0",
		},
	)

	return
}

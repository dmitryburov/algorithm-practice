package main

import (
	"fmt"
	"strings"
	"testing"
)

type task struct {
	K, Result int
	Matrix    [][]string
}

func TestTask(t *testing.T) {

	var matrixCnt = 4
	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Example %d", 0+1), func(t *testing.T) {
			if res := solution(matrixCnt, taskItems[i].K, taskItems[i].Matrix); res != taskItems[i].Result {
				t.Errorf("Неверный ответ решения! Ответ: %d. Правильный: %d", res, taskItems[i].Result)
			}
		})
	}
}

// generateTasks создает задачи для теста
func generateTasks() (tasks []task) {
	tasks = append(
		tasks,
		task{
			K:      3,
			Result: 2,
			Matrix: generateMatrix([]string{"1231", "2..2", "2..2", "2..2"}),
		},
		task{
			K:      4,
			Result: 1,
			Matrix: generateMatrix([]string{"1111", "9999", "1111", "9911"}),
		},
		task{
			K:      4,
			Result: 0,
			Matrix: generateMatrix([]string{"1111", "1111", "1111", "1111"}),
		},
	)

	return
}

// generateMatrix генерирует матрицу для задачи
func generateMatrix(matrix []string) (result [][]string) {
	for i := 0; i < len(matrix); i++ {
		result = append(result, strings.Split(matrix[i], ""))
	}

	return
}

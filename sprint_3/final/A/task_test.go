package main

import (
	"fmt"
	"testing"
)

type task struct {
	K      int
	Input  []int
	Result int
}

func TestTask(t *testing.T) {
	tasks := generateTasks()
	for i := 0; i < len(tasks); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := brokenSearch(tasks[i].Input, tasks[i].K)
			if res != tasks[i].Result {
				t.Errorf("Неверный ответ решения!\nОтвет: %d \nВерно: %d", res, tasks[i].Result)
			}
		})
	}
}

func generateTasks() []task {
	return []task{
		{
			K:      5,
			Input:  []int{19, 21, 100, 101, 1, 4, 5, 7, 12},
			Result: 6,
		},
		{
			K:      1,
			Input:  []int{5, 1},
			Result: 1,
		},
		{
			K:      5,
			Input:  []int{0, 2, 6, 7, 8, 9, 10},
			Result: -1,
		},
		{
			K:      25,
			Input:  []int{3271, 3298, 3331, 3397, 3407, 3524, 3584, 3632, 3734, 3797, 3942, 4000, 4180, 4437, 4464, 4481, 4525, 4608, 4645, 4803, 4804, 4884, 4931, 4965, 5017, 5391, 5453, 5472, 5671, 5681, 5959, 6045, 6058, 6301, 6529, 6621, 6961, 7219, 7291, 7372, 7425, 7517, 7600, 7731, 7827, 7844, 7987, 8158, 8169, 8265, 8353, 8519, 8551, 8588, 8635, 9209, 9301, 9308, 9336, 9375, 9422, 9586, 9620, 9752, 9776, 9845, 9906, 9918, 16, 25, 45, 152, 199, 309, 423, 614, 644, 678, 681, 725, 825, 830, 936, 1110, 1333, 1413, 1617, 1895, 1938, 2107, 2144, 2184, 2490, 2517, 2769, 2897, 2970, 3023, 3112, 3156},
			Result: 69,
		},
	}
}

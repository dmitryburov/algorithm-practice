package main

import (
	"fmt"
	"strings"
	"testing"
)

type task struct {
	N      int
	Result string
}

func TestTask(t *testing.T) {

	var taskItems = generateTasks()

	for i := 0; i < len(taskItems); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			s := strings.Builder{}
			solution(taskItems[i].N, "", 0, 0, &s)
			if strings.TrimRight(s.String(), "\n") != taskItems[i].Result {
				t.Errorf("Неверный ответ решения!\nОтвет: %s \nВерно: %s", s.String(), taskItems[i].Result)
			}
		})
	}
}

func generateTasks() (tasks []task) {
	tasks = append(
		tasks,
		task{
			N: 3,
			Result: `((()))
(()())
(())()
()(())
()()()`},
		task{
			N: 2,
			Result: `(())
()()`},
	)

	return
}

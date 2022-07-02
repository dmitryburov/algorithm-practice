package main

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

type TestItem struct {
	input  io.Reader
	output string
}

func TestTask(t *testing.T) {
	tasks := generateTasks()

	for i, v := range tasks {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			buf := strings.Builder{}
			Solution(v.input, &buf)
			if buf.String() != v.output {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%v \nВерно: \n%v", buf.String(), v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{strings.NewReader(`5 3
3 2
3 4
2 5
`), "1 3 2 4 5"},
		{strings.NewReader(`6 3
6 4
4 1
5 1
`), "2 3 5 6 4 1"},
		{strings.NewReader(`4 0
`), "1 2 3 4"},
	}
}

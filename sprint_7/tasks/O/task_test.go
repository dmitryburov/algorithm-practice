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
		{strings.NewReader(`3 3
1 2
1 2
2 3
1 3
`), "2"},
		{strings.NewReader(`5 3
1 2
3 4
4 5
1 5
`), "0"},
		{strings.NewReader(`3 3
1 2
2 3
1 3
1 1
`), "1"},
	}
}

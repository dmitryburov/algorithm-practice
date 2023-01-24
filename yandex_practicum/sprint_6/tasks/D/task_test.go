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
			if strings.TrimSpace(buf.String()) != v.output {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%v \nВерно: \n%v", buf.String(), v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{strings.NewReader(`4 4
1 2
2 3
3 4
1 4
3
`), "3 2 4 1"},
		{strings.NewReader(`2 1
1 2
1
`), "1 2"},
		{strings.NewReader(`1 0
1
`), "1"},
		{strings.NewReader(`6 7
3 2
5 4
3 1
1 4
1 6
1 2
1 5
1
`), "1 2 3 4 5 6"},
		{strings.NewReader(`6 8
5 6
1 4
2 5
2 4
6 4
5 4
1 5
3 1
1
`), "1 3 4 5 2 6"},
	}
}

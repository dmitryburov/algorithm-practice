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
		{strings.NewReader(`4 4
1 2 5
1 3 6
2 4 8
3 4 3
`), "19"},
		{strings.NewReader(`3 3
1 2 1
1 2 2
2 3 1
`), "3"},
		{strings.NewReader(`2 0
`), "Oops! I did it again"},
		{strings.NewReader(`1 0
`), "0"},
		{strings.NewReader(`10 20
9 10 4
2 2 4
4 2 8
10 5 3
1 10 6
7 4 2
10 10 6
3 7 4
8 9 4
8 10 7
6 10 10
2 8 8
3 8 1
3 10 3
9 5 8
10 10 2
1 8 1
10 1 5
3 6 10
9 10 8
`), "69"}}
}

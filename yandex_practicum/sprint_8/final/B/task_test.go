package main

import (
	"bufio"
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
			s, n := initInput(bufio.NewScanner(v.input))
			Solution(s, n, &buf)
			if buf.String() != v.output {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%v \nВерно: \n%v", buf.String(), v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{strings.NewReader(`examiwillpasstheexam
5
will
pass
the
exam
i
`), "YES"},
		{strings.NewReader(`abacaba
2
abac
caba
`), "NO"},
		{strings.NewReader(`abacaba
3
abac
caba
aba
`), "YES"},
	}
}

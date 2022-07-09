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
		{strings.NewReader(`5
caba
aba
caba
abac
aba
`), "aba"},
		{strings.NewReader(`3
b
bc
bcd
`), "b"},
		{strings.NewReader(`10
ciwlaxtnhhrnenw
ciwnvsuni
ciwaxeujmsmvpojqjkxk
ciwnvsuni
ciwnvsuni
ciwuxlkecnofovq
ciwuxlkecnofovq
ciwodramivid
ciwlaxtnhhrnenw
ciwnvsuni
`), "ciwnvsuni"},
	}
}

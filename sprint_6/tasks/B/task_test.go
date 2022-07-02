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
			if strings.Compare(buf.String(), v.output) == 0 {
				t.Errorf("Неверный ответ решения!\nОтвет: \n%v \nВерно: \n%v", buf.String(), v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{strings.NewReader(`5 3
1 3
2 3
5 2
`), `0 0 1 0 0 
0 0 1 0 0 
0 0 0 0 0 
0 0 0 0 0 
0 1 0 0 0 
`},
		{strings.NewReader(`10 55
1 3
1 5
1 6
1 9
1 10
2 5
2 7
2 8
2 10
3 1
3 6
3 7
4 1
4 2
4 3
4 5
4 7
4 8
4 10
5 3
5 4
5 6
5 7
5 8
5 9
5 10
6 1
6 2
6 3
6 5
6 8
6 9
6 10
7 1
7 3
7 8
8 1
8 2
8 3
8 4
8 5
8 6
8 7
8 9
9 2
9 4
9 5
9 6
9 8
10 1
10 2
10 4
10 5
10 6
10 7
`), `0 0 1 0 1 1 0 0 1 1 
0 0 0 0 1 0 1 1 0 1 
1 0 0 0 0 1 1 0 0 0 
1 1 1 0 1 0 1 1 0 1 
0 0 1 1 0 1 1 1 1 1 
1 1 1 0 1 0 0 1 1 1 
1 0 1 0 0 0 0 1 0 0 
1 1 1 1 1 1 1 0 1 0 
0 1 0 1 1 1 0 1 0 0 
1 1 0 1 1 1 1 0 0 0 
`},
	}
}

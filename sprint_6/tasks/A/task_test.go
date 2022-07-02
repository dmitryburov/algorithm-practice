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
			GetList(v.input, &buf)
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
5 2`), `1 3 
1 3 
0 
0 
1 2`},
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
10 7`), `5 3 5 6 9 10
4 5 7 8 10
3 1 6 7
7 1 2 3 5 7 8 10
7 3 4 6 7 8 9 10
7 1 2 3 5 8 9 10
3 1 3 8
8 1 2 3 4 5 6 7 9
5 2 4 5 6 8
6 1 2 4 5 6 7`},
		{strings.NewReader(`15 45
1 4
1 11
1 14
2 1
2 5
2 7
3 2
3 5
3 7
3 8
3 9
4 9
5 2
5 4
5 7
5 12
6 9
6 11
6 15
7 6
8 3
8 4
8 11
8 12
8 15
9 3
10 2
10 8
11 2
11 15
12 5
12 8
13 3
13 5
13 7
13 9
13 10
13 15
14 1
14 2
14 3
14 5
14 13
15 6
15 8`), `3 4 11 14
3 1 5 7
5 2 5 7 8 9
1 9
4 2 4 7 12
3 9 11 15
1 6
5 3 4 11 12 15
1 3
2 2 8
2 2 15
2 5 8
6 3 5 7 9 10 15
5 1 2 3 5 13
2 6 8`},
	}
}

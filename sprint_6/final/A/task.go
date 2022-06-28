package main

/*
С этой задачей пришлось возвращаться к теории и перечитывать.

--- Посылка


--- Принцип работы


--- Временная сложность
n - количество вершин в графе
m - количество ребер в графе

Сложность составляет O(m*log(n)).

--- Пространственная сложность
Хранятся пары вершин и ребра, следовательно сложность будет  O(n+m)
*/

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const ErrOops = "Oops! I did it again"

func main() {
	res := strings.Builder{}
	Solution(os.Stdin, &res)

	fmt.Println(res.String())
}

// Solution решение задачи
func Solution(r io.Reader, w *strings.Builder) {
	var err error
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	// количество вершин и ребер
	var n, m int
	nmData := strings.Fields(scanner.Text())
	n, err = strconv.Atoi(nmData[0])
	if err != nil {
		log.Fatal(err)
	}
	m, err = strconv.Atoi(nmData[1])
	if err != nil {
		log.Fatal(err)
	}

	w.WriteString(ErrOops)
}

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
func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	// количество вершин и ребер
	var n, m int
	nmData := strings.Fields(scanner.Text())
	n, _ = strconv.Atoi(nmData[0])
	m, _ = strconv.Atoi(nmData[1])

	// fix empty
	if n > 1 && m == 0 {
		s.WriteString(ErrOops)
		return
	}

	edgesList := make([][]int, n)

	var v, u, w int
	for i := 1; i <= m; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		v, _ = strconv.Atoi(edgeData[0])
		u, _ = strconv.Atoi(edgeData[1])
		m, _ = strconv.Atoi(edgeData[2])

		// fix если равны
		if v == u {
			continue
		}

		//if _, ok := edgesList[u][v]; !ok || (ok && edgesList[u][v] < w) {
		//	edgesList[v][u] = w
		//	edgesList[u][v] = w
		//}
	}

	var mst int
	var edges []int

	res, err := FindMST(edgesList, edges, mst)
	if err != nil {
		s.WriteString(ErrOops)
		return
	}

	s.WriteString(fmt.Sprint(res))
}

// FindMST поиск минимального остовного дерева
func FindMST(graph [][]int, edges []int, mst int) (int, error) {

	return 0, nil
}

// AddPeak добавление вершины
func AddPeak() {

}

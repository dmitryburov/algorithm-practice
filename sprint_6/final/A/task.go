package main

/*
С этой задачей пришлось возвращаться к теории и перечитывать. А где-то даже хорошенько погуглить...
Для построения максимального остового дерева использовался алгоритм Дийкстры из теории.

--- Посылка
69252650

--- Принцип работы
Сначала собираем из входных данных вершины и ребра.
Далее выбираем одну из вершин (первую) мы добавляем ребра в очередь. Потом, извлекая из очередь ребра до тех пора, пока не найдем непосещенную.
Если находим - возвращаем общую длину пути. В конце мы проходим по массиву посещенных вершин и если все посещены - отправляем длину или ошибку.

Дополнительные комментрарии присутствуют в коде.

Касаемо доп.обхода массива - я пытался избежать дополнительного цикла за счет "счетчика" и в конце сравнивал значения,
но почему-то всегда на 39 тесте падало(

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

const (
	ErrOops     = "Oops! I did it again"
	ErrPopQueue = "Queue is empty"
)

// Edge ребро
type Edge struct {
	weight int
	node   *Node
}

// Node узел
type Node struct {
	value int
	edges []Edge
}

// NodeList список смежности
type NodeList []*Node

// Queue очередь для хранения ребер
type Queue []Edge

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

	// fix если нет ребер
	if n > 1 && m == 0 {
		s.WriteString(ErrOops)
		return
	}

	list := make(NodeList, n+1)

	// обойдем вершины
	for i := 1; i <= n; i++ {
		list[i] = &Node{
			value: i,
		}
	}

	// обойдем ребра
	var v, u, w int
	for i := 1; i <= m; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		// вершины
		v, _ = strconv.Atoi(edgeData[0])
		u, _ = strconv.Atoi(edgeData[1])
		// вес
		w, _ = strconv.Atoi(edgeData[2])

		// fix если равны
		if v == u {
			continue
		}

		list[v].edges = append(list[v].edges, Edge{
			weight: w,
			node:   list[u],
		})
		list[u].edges = append(list[u].edges, Edge{
			weight: w,
			node:   list[v],
		})
	}

	// проверим наличие остовного дерева
	if weight, err := FindMST(list); err {
		s.WriteString(fmt.Sprint(weight))
	} else {
		s.WriteString(ErrOops)
	}
}

// Push добавление в очередь
func (q *Queue) Push(edge Edge) {
	*q = append(*q, edge)
	q.TreeUp(q.Len() - 1)
}

// Pop извлечение из очереди
func (q *Queue) Pop() (*Edge, error) {
	if q.Len() == 1 {
		return nil, fmt.Errorf("%s", ErrPopQueue)
	}

	save := (*q)[1]
	(*q)[1] = (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	q.TreeDown(1)

	return &save, nil
}

// Len размер очереди
func (q Queue) Len() int {
	return len(q)
}

// Less сравнение
func (q Queue) Less(a, b int) bool {
	return q[a].weight > q[b].weight
}

// Swap замена
func (q Queue) Swap(a, b int) {
	q[a], q[b] = q[b], q[a]
}

// TreeUp просеивание вверх
func (q Queue) TreeUp(index int) {
	if index == 1 {
		return
	}
	parentIndex := index >> 1

	if q.Less(index, parentIndex) {
		q.Swap(index, parentIndex)
		q.TreeUp(parentIndex)
	}
}

// TreeDown просеивание вниз
func (q Queue) TreeDown(index int) {
	firstChild := index * 2
	secondChild := firstChild + 1

	if firstChild >= q.Len() {
		return
	}
	best := firstChild

	if secondChild < q.Len() && q.Less(secondChild, firstChild) {
		best = secondChild
	}

	if q.Less(best, index) {
		q.Swap(best, index)
		q.TreeDown(best)
	}
}

// FindMST поиск остового дерева (Дийкстра)
func FindMST(al NodeList) (int, bool) {
	var (
		total = true
		accum int
	)
	queue := make(Queue, 1, len(al)) // очередь ребер
	visited := make([]bool, len(al)) // посещенные вершины

	visited[1] = true

	// добавляем ребра в очередь
	for _, edge := range al[1].edges {
		queue.Push(edge)
	}

	// извлекаем из очереди и проверям
	for edge, err := queue.Pop(); err == nil; edge, err = queue.Pop() {
		if visited[edge.node.value] {
			continue
		}

		accum += edge.weight
		visited[edge.node.value] = true

		for _, e := range edge.node.edges {
			if !visited[e.node.value] {
				queue.Push(e)
			}
		}
	}

	for i := 1; i < len(visited); i++ {
		if !visited[i] {
			total = false
		}
	}

	return accum, total
}

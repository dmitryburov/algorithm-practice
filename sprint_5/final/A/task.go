package main

/*
--- ID посылки:
69190328

--- Принцип работы:
Благодаря занятиям и методам сортировок из предыдущих спринтов + достаточно хорошо описаная теория о просеивании, то можно сказать, что нужно было
просто реализовать.

Если кратко, то Append добавляет эелемент в очередь, не забывая чтобы в самом начале был самый приоритетный элемент,
с помощью UpSifting
После того, как все эелементы добавлены, забираются приоритетные элементы с помощью Next, который
поддерживает, чтобы самый второй по приоритетности элемент оказался на вершине кучи с помощью DownSifting

Методы дополнительно прокомментировал в коде.

--- Временная сложность:
В целом временная сложность равна O(n * log n), но почему-то меня терзают сомнения, возможно не учел константу где-то?

--- Пространственная сложность:
Ну если отсекать входные данные, то по сути мы только используем затраты на рекурсии при вставке/получении элемента, которые зависят от высоты дерева.
В рекурсиях мы не выделяем память на доп структуры, а работаем с текущей, следовательно пространственная будет O(log n)
*/
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// CompetitionResult структура элемента
type CompetitionResult struct {
	name     string
	complete int
	fail     int
}

// Queue структура очереди
type Queue struct {
	queue     []CompetitionResult
	lastIndex int
}

func main() {
	res := strings.Builder{}
	PyramidSort(os.Stdin, &res)

	fmt.Println(res.String())
}

// PyramidSort основная функция на вход
func PyramidSort(r io.Reader, w *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	pq := Queue{queue: make([]CompetitionResult, n+1), lastIndex: 0}

	for i := 0; i < n; i++ {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		name := fields[0]
		complete, _ := strconv.Atoi(fields[1])
		fail, _ := strconv.Atoi(fields[2])
		pq.Append(CompetitionResult{name, complete, fail})
	}

	for i := 0; i < n; i++ {
		w.WriteString(pq.Next().name)
		w.WriteByte('\n')
	}
}

// Append добавление элемента в очередь
func (pq *Queue) Append(cr CompetitionResult) {
	pq.lastIndex++
	pq.queue[pq.lastIndex] = cr
	pq.UpSifting(pq.lastIndex)
}

// Next получение приоритетного элемента из кучи
func (pq *Queue) Next() CompetitionResult {
	defer func() {
		pq.queue[1] = pq.queue[pq.lastIndex]
		pq.queue = pq.queue[:pq.lastIndex]
		pq.DownSifting(1)
		pq.lastIndex--
	}()

	return pq.queue[1]
}

// UpSifting просеивания вверх
func (pq Queue) UpSifting(index int) {
	if index == 1 {
		return
	}

	parentIndex := index >> 1

	if pq.Less(index, parentIndex) {
		pq.Swap(index, parentIndex)
		pq.UpSifting(parentIndex)
	}
}

// DownSifting просеивания вниз
func (pq Queue) DownSifting(index int) {
	left := index * 2
	right := index*2 + 1
	if left >= pq.Len() {
		return
	}

	best := 0

	if right < pq.Len() && pq.Less(right, left) {
		best = right
	} else {
		best = left
	}

	if pq.Less(best, index) {
		pq.Swap(best, index)
		pq.DownSifting(best)
	}
}

// Less проверка приоритета
func (pq Queue) Less(a, b int) bool {
	if pq.queue[a].complete == pq.queue[b].complete {
		if pq.queue[a].fail == pq.queue[b].fail {
			return pq.queue[a].name < pq.queue[b].name
		} else {
			return pq.queue[a].fail <= pq.queue[b].fail
		}
	} else {
		return pq.queue[a].complete >= pq.queue[b].complete
	}
}

// Swap классическая функция смены
func (pq Queue) Swap(a, b int) {
	pq.queue[a], pq.queue[b] = pq.queue[b], pq.queue[a]
}

// Len текущий размер очереди
func (pq Queue) Len() int {
	return len(pq.queue)
}

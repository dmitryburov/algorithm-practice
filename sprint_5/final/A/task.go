package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type CompetitionResult struct {
	name       string
	completeEx int
	penalty    int
}

type PriorityQueue struct {
	queue         []CompetitionResult
	lastItemIndex int
}

func (pq *PriorityQueue) Append(cr CompetitionResult) {
	pq.lastItemIndex++
	pq.queue[pq.lastItemIndex] = cr
	pq.SiftUpBalance(pq.lastItemIndex)
}

func (pq *PriorityQueue) Next() CompetitionResult {

	defer func() {
		pq.queue[1] = pq.queue[pq.lastItemIndex]
		pq.queue = pq.queue[:pq.lastItemIndex]
		pq.SiftDownBalance(1)
		pq.lastItemIndex--
	}()
	return pq.queue[1]
}

func (pq PriorityQueue) SiftUpBalance(index int) {
	if index == 1 {
		return
	}
	parrentIndex := index >> 1
	if pq.Less(index, parrentIndex) {
		pq.Swap(index, parrentIndex)
		pq.SiftUpBalance(parrentIndex)
	}

}

func (pq PriorityQueue) SiftDownBalance(index int) {
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
		pq.SiftDownBalance(best)
	}
}

func (pq PriorityQueue) Less(a, b int) bool {
	if pq.queue[a].completeEx == pq.queue[b].completeEx {
		if pq.queue[a].penalty == pq.queue[b].penalty {
			return compareStings(pq.queue[a].name, pq.queue[b].name)
		} else {
			return pq.queue[a].penalty <= pq.queue[b].penalty
		}
	} else {
		return pq.queue[a].completeEx >= pq.queue[b].completeEx
	}
}

func (pq PriorityQueue) Swap(a, b int) {
	pq.queue[a], pq.queue[b] = pq.queue[b], pq.queue[a]
}

func (pq PriorityQueue) Len() int {
	return len(pq.queue)
}

func PyramidSort(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	pq := PriorityQueue{queue: make([]CompetitionResult, n+1), lastItemIndex: 0}

	for i := 0; i < n; i++ {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		name := fields[0]
		completeEx, _ := strconv.Atoi(fields[1])
		penalty, _ := strconv.Atoi(fields[2])
		pq.Append(CompetitionResult{name, completeEx, penalty})
	}

	writer := strings.Builder{}
	for i := 0; i < n; i++ {
		writer.WriteString(pq.Next().name)
		writer.WriteByte('\n')
	}

	fmt.Println(writer.String())
}

func compareStings(s1, s2 string) bool {
	s1Lenght := len(s1)
	s2Lenght := len(s2)
	minLenght := 0
	if s1Lenght > s2Lenght {
		minLenght = s2Lenght
	} else {
		minLenght = s1Lenght
	}
	i, sum1, sum2 := 0, 0, 0
	for i < minLenght && sum1 == sum2 {
		sum1 += int(s1[i])
		sum2 += int(s2[i])
		i++
	}
	if sum1 == sum2 {
		return s1Lenght < s2Lenght
	}
	return sum1 < sum2
}

func main() {
	PyramidSort(os.Stdin)
}

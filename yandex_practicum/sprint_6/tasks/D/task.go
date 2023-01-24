package main

// TODO переделать
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value    int
	PointsTo AdjacencyList
}

type AdjacencyList []*Node

type Queue struct {
	CurrentSize  int
	Queue        AdjacencyList
	FirstInQueue int
	EmptyIndex   int
}

func (q *Queue) Push(node *Node) {
	q.Queue[q.EmptyIndex] = node
	q.EmptyIndex++
	q.EmptyIndex %= q.CurrentSize
	predictedIndex := (q.EmptyIndex + 1) % q.CurrentSize
	if q.FirstInQueue == predictedIndex {
		q.Rebalance()
	}
}

func (q *Queue) Pop() (*Node, error) {
	if q.FirstInQueue == q.EmptyIndex {
		return nil, errors.New("the queue is empty")
	}
	// тут я бы использовал defer, но это довольно нагруженная часть при ребалансировании, так что без него
	save := q.Queue[q.FirstInQueue]
	q.FirstInQueue++
	q.FirstInQueue %= q.CurrentSize
	return save, nil
}

func (q *Queue) Rebalance() {
	newQueue := make(AdjacencyList, q.CurrentSize*2)
	newQueueIndex := 0
	for i, err := q.Pop(); err == nil; i, err = q.Pop() {
		newQueue[newQueueIndex] = i
		newQueueIndex++
	}
	q.Queue = newQueue
	q.CurrentSize *= 2
	q.FirstInQueue = 0
	q.EmptyIndex = newQueueIndex
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())
	peaks, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	AL := make(AdjacencyList, peaks+1)
	for i := 0; i < peaks; i++ {
		currentIndex := i + 1
		AL[currentIndex] = &Node{
			Value:    currentIndex,
			PointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		peakA, _ := strconv.Atoi(edgeData[0])
		peakB, _ := strconv.Atoi(edgeData[1])
		AL[peakA].PointsTo = append(AL[peakA].PointsTo, AL[peakB])
		AL[peakB].PointsTo = append(AL[peakB].PointsTo, AL[peakA])
	}
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	queue := Queue{
		CurrentSize:  16,
		Queue:        make(AdjacencyList, 16),
		FirstInQueue: 0,
		EmptyIndex:   0,
	}

	queue.Push(AL[N])
	colors := make([]string, peaks+1)

	colors[AL[N].Value] = "gray"
	bfs(&queue, colors, s)
}

func bfs(q *Queue, colors []string, writer io.StringWriter) {
	node, err := q.Pop()
	if err != nil {
		return
	}
	node.PointsTo = merge(node.PointsTo)
	_, err = writer.WriteString(strconv.Itoa(node.Value) + " ")
	if err != nil {
		return
	}
	for _, n := range node.PointsTo {
		if colors[n.Value] == "" {
			colors[n.Value] = "gray"
			q.Push(n)
		}
	}
	colors[node.Value] = "black"
	bfs(q, colors, writer)
}

func merge(list AdjacencyList) AdjacencyList {
	listLength := len(list)
	if listLength < 2 {
		return list
	}
	mid := listLength >> 1
	aList := merge(list[:mid])
	aLength := len(aList)
	bList := merge(list[mid:])
	bLength := len(bList)
	aIndex := 0
	bIndex := 0
	result := make(AdjacencyList, 0, aLength+bLength)
	for aIndex < aLength || bIndex < bLength {
		if aIndex == aLength {
			result = append(result, bList[bIndex])
			bIndex++
			continue
		}
		if bIndex == bLength {
			result = append(result, aList[aIndex])
			aIndex++
			continue
		}
		if aList[aIndex].Value < bList[bIndex].Value {
			result = append(result, aList[aIndex])
			aIndex++
		} else {
			result = append(result, bList[bIndex])
			bIndex++
		}
	}
	return result
}

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(strings.TrimSpace(s.String()))
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	to    []*Node
}

type Stack struct {
	cnt   int
	stack []string
}

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	initData := strings.Fields(scanner.Text())
	peaks, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])

	list := make([]*Node, peaks+1)
	for i := 1; i <= peaks; i++ {
		list[i] = &Node{
			value: i,
		}
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		data := strings.Fields(scanner.Text())
		peakA, _ := strconv.Atoi(data[0])
		peakB, _ := strconv.Atoi(data[1])
		list[peakA].to = append(list[peakA].to, list[peakB])
	}

	colors := make([]string, peaks+1)
	stack := Stack{
		cnt:   peaks - 1,
		stack: make([]string, peaks),
	}

	for i := peaks; i > 0; i-- {
		if colors[i] == "" {
			Sort(list[i], colors, &stack)
		}
	}

	s.WriteString(fmt.Sprint(&stack))
}

func (s *Stack) String() string {
	return strings.Join(s.stack, " ")
}

func (s *Stack) Push(v string) {
	s.stack[s.cnt] = v
	s.cnt--
}

func Sort(n *Node, colors []string, stack *Stack) {
	colors[n.value] = "gray"
	for _, node := range n.to {
		if colors[node.value] == "" {
			Sort(node, colors, stack)
		}
	}
	colors[n.value] = "black"
	stack.Push(strconv.Itoa(n.value))
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	SUCCESS = "YES"
	FAIL    = "NO"
)

type Node struct {
	value int
	to    NodeList
}

type NodeList []*Node

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	initData := strings.Fields(scanner.Text())

	vertexes, _ := strconv.Atoi(initData[0])
	edges, _ := strconv.Atoi(initData[1])
	list := make(NodeList, vertexes+1)

	for i := 1; i <= vertexes; i++ {
		list[i] = &Node{
			value: i,
			to:    nil,
		}
	}

	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		a, _ := strconv.Atoi(edgeData[0])
		b, _ := strconv.Atoi(edgeData[1])
		if a != b {
			list[a].to = append(list[a].to, list[b])
			list[b].to = append(list[b].to, list[a])
		}
	}

	for _, node := range list {
		if node == nil {
			continue
		}
		visited := make([]bool, vertexes+1)
		for _, n := range node.to {
			visited[n.value] = true
		}
		for i, isVisited := range visited {
			if i == 0 || i == node.value {
				continue
			}
			if !isVisited {
				s.WriteString(FAIL)
				return
			}
		}
	}

	s.WriteString(SUCCESS)
}

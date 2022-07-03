package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Color int

const (
	_ Color = iota
	WHITE
	GRAY
	BLACK
)

type Node struct {
	to    NodeList
	value int
}

type NodeList []*Node

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(strings.TrimSpace(s.String()))
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	var (
		n, m int
		err  error
	)

	values := strings.Fields(scanner.Text())
	n, err = strconv.Atoi(values[0])
	if err != nil {
		log.Fatalln(err)
	}

	m, err = strconv.Atoi(values[1])
	if err != nil {
		log.Fatalln(err)
	}

	matrix := make(NodeList, n+1)
	colors := make([]Color, n+1)

	for i := 0; i < m; i++ {
		scanner.Scan()
		edge := strings.Fields(scanner.Text())

		a, err := strconv.Atoi(edge[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err := strconv.Atoi(edge[1])
		if err != nil {
			log.Fatal(err)
		}

		if matrix[a] == nil {
			colors[a] = WHITE
			matrix[a] = &Node{value: a}
		}
		if matrix[b] == nil {
			colors[b] = WHITE
			matrix[b] = &Node{value: b}
		}
		matrix[a].to = append(matrix[a].to, matrix[b])
		matrix[b].to = append(matrix[b].to, matrix[a])
	}

	scanner.Scan()
	start, _ := strconv.Atoi(scanner.Text())
	DFS(matrix[start], colors, s)
}

func DFS(node *Node, colors []Color, s *strings.Builder) {
	if node == nil || colors[node.value] == BLACK {
		return
	}

	colors[node.value] = BLACK
	s.WriteString(fmt.Sprint(node.value) + " ")

	for _, v := range node.to {
		if colors[v.value] != BLACK {
			DFS(v, colors, s)
		}
	}
}

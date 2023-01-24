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
	value    int
	pointsTo AdjacencyList
}

type AdjacencyList []*Node

const (
	RED     = "red"
	BLUE    = "blue"
	SUCCESS = "YES"
	FAIL    = "NO"
)

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
	AL := make(AdjacencyList, vertexes+1)
	for i := 0; i < vertexes; i++ {
		correctIndex := i + 1
		AL[correctIndex] = &Node{
			value:    correctIndex,
			pointsTo: nil,
		}
	}
	for i := 0; i < edges; i++ {
		scanner.Scan()
		edgeData := strings.Fields(scanner.Text())
		vertexA, _ := strconv.Atoi(edgeData[0])
		vertexB, _ := strconv.Atoi(edgeData[1])
		AL[vertexA].pointsTo = append(AL[vertexA].pointsTo, AL[vertexB])
		AL[vertexB].pointsTo = append(AL[vertexB].pointsTo, AL[vertexA])
	}
	colors := make([]string, vertexes+1)
	for _, node := range AL {
		if node == nil {
			continue
		}
		if colors[node.value] == "" {
			ok := dfs(node, colors, BLUE)
			if !ok {
				s.WriteString(FAIL)
				return
			}
		}
	}

	s.WriteString(SUCCESS)
}

func dfs(node *Node, colors []string, currentColor string) bool {
	colors[node.value] = currentColor
	OC := color(currentColor)
	for _, n := range node.pointsTo {
		if colors[n.value] == "" {
			ok := dfs(n, colors, OC)
			if !ok {
				return false
			}
		}
		if colors[n.value] == currentColor {
			return false
		}
	}
	return true
}

func color(color string) string {
	if color == BLUE {
		return RED
	}

	return BLUE
}

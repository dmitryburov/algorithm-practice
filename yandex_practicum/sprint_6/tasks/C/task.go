package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Vertex struct {
	node  int
	color Color
}

type Color int

const (
	WHITE Color = iota
	GRAY
	BLACK
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	var n, m int
	var err error

	values := strings.Fields(scanner.Text())
	n, err = strconv.Atoi(values[0])
	if err != nil {
		log.Fatalln(err)
	}

	m, err = strconv.Atoi(values[1])
	if err != nil {
		log.Fatalln(err)
	}

	var a, b int
	var data = make([]int, m*2)
	for i := 0; i < len(data); i++ {
		scanner.Scan()
		edge := strings.Fields(scanner.Text())

		a, err = strconv.Atoi(edge[0])
		if err != nil {
			log.Fatal(err)
		}
		data[i] = a
		i++

		b, err = strconv.Atoi(edge[1])
		if err != nil {
			log.Fatal(err)
		}
		data[i] = b
	}

	scanner.Scan()
	start, _ := strconv.Atoi(scanner.Text())

	vertexex := Vertexes(data, n, m)
	colors := make([]Color, len(vertexex))

	result := DFS(colors, make([]int, 0), vertexex, start)
	s.WriteString(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]"))
}

func Vertexes(data []int, vertexCount, edgeCount int) [][]Vertex {
	vertexCount++

	vertexes := make([][]Vertex, vertexCount)
	for i := range vertexes {
		vertexes[i] = make([]Vertex, 0)
	}

	for i := 0; i < edgeCount*2; i += 2 {
		from, to := data[i], data[i+1]
		vertexes[from] = append(vertexes[from], Vertex{
			node: to,
		})

		vertexes[to] = append(vertexes[to], Vertex{
			node: from,
		})
	}

	for i := range vertexes {
		sort.Slice(vertexes[i], func(first, second int) bool {
			return vertexes[i][first].node < vertexes[i][second].node
		})
	}

	return vertexes
}

func DFS(colorArray []Color, prevResult []int, vertexArray [][]Vertex, currentWay int) []int {
	colorArray[currentWay] = GRAY
	ways := vertexArray[currentWay]
	prevResult = append(prevResult, currentWay)

	for _, way := range ways {
		if colorArray[way.node] == WHITE {
			prevResult = DFS(colorArray, prevResult, vertexArray, way.node)
		}
	}

	colorArray[currentWay] = BLACK

	return prevResult
}

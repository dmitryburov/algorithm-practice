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

const MOD = 1000000007

type Vertex struct {
	value    int
	pointsTo []*Vertex
}

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	var err error
	var n, m int
	var arrStr []string

	scanner.Scan()
	arrStr = strings.Fields(scanner.Text())
	n, err = strconv.Atoi(arrStr[0])
	if err != nil {
		log.Fatal(err)
	}

	m, err = strconv.Atoi(arrStr[1])
	if err != nil {
		log.Fatal(err)
	}

	list := make([]Vertex, n+1)
	for i := 1; i < n+1; i++ {
		list[i].value = i
	}

	var a, b int
	for i := 0; i < m; i++ {
		scanner.Scan()
		arrStr = strings.Fields(scanner.Text())

		a, err = strconv.Atoi(arrStr[0])
		if err != nil {
			log.Fatal(err)
		}

		b, err = strconv.Atoi(arrStr[1])
		if err != nil {
			log.Fatal(err)
		}

		list[a].pointsTo = append(list[a].pointsTo, &list[b])
	}

	scanner.Scan()
	arrStr = strings.Fields(scanner.Text())
	from, _ := strconv.Atoi(arrStr[0])
	to, _ := strconv.Atoi(arrStr[1])

	paths := make([]int, n+1)
	visited := make([]bool, n+1)
	paths[to] = 1

	dfs(&list[from], visited, paths)
	s.WriteString(fmt.Sprint(paths[from]))
}

func dfs(vertex *Vertex, visited []bool, paths []int) {
	visited[vertex.value] = true
	for _, v := range vertex.pointsTo {
		if !visited[v.value] {
			dfs(v, visited, paths)
		}
		paths[vertex.value] = (paths[vertex.value] + paths[v.value]) % MOD
	}
}

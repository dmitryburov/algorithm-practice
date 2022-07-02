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

type NodeList []*Node

type Node struct {
	value  int
	points []*Node
}

func main() {
	s := strings.Builder{}
	GetList(os.Stdin, &s)

	fmt.Println(s.String())
}

func GetList(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	var n, m int
	var err error

	nValues := strings.Fields(scanner.Text())
	n, err = strconv.Atoi(nValues[0])
	if err != nil {
		log.Fatalln(err)
	}
	m, err = strconv.Atoi(nValues[1])
	if err != nil {
		log.Fatalln(err)
	}

	list := make(NodeList, n+1)

	for i := 0; i < m; i++ {
		scanner.Scan()
		edge := strings.Fields(scanner.Text())

		var u, v int
		u, err = strconv.Atoi(edge[0])
		if err != nil {
			log.Fatalln(err)
		}
		v, err = strconv.Atoi(edge[1])
		if err != nil {
			log.Fatalln(err)
		}

		if list[u] == nil {
			list[u] = &Node{value: u}
		}

		if list[v] == nil {
			list[v] = &Node{value: v}
		}

		list[u].points = append(list[u].points, list[v])
	}

	for i, v := range list {
		if i == 0 {
			continue
		}

		if v == nil {
			s.WriteString("0\n")
			continue
		}

		if len(v.points) == 0 {
			s.WriteString("0\n")
			continue
		}

		s.WriteString(strconv.Itoa(len(v.points)) + " ")

		for _, g := range v.points {
			s.WriteString(strconv.Itoa(g.value) + " ")
		}

		if i != len(list)-1 {
			s.WriteByte('\n')
		}
	}
}

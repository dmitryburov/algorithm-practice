package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Matrix [][]int

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	nValues := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(nValues[0])
	m, _ := strconv.Atoi(nValues[1])

	matrix := make(Matrix, n)
	for i := 0; i < m; i++ {
		scanner.Scan()
		edge := strings.Fields(scanner.Text())
		edgeA, _ := strconv.Atoi(edge[0])
		edgeB, _ := strconv.Atoi(edge[1])
		edgeA--
		edgeB--
		if matrix[edgeA] == nil {
			matrix[edgeA] = make([]int, n)
		}
		matrix[edgeA][edgeB] = 1
	}
	for _, v := range matrix {
		if v == nil {
			v = make([]int, n)
		}
		for _, vV := range v {
			s.WriteString(fmt.Sprint(vV) + " ")
		}
		s.WriteString("\n")
	}
}

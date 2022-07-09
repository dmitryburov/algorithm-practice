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

type Matrix [][]int

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

	matrix := make(Matrix, n+1)
	dp := make(Matrix, n+2)

	for i := 0; i < n+1; i++ {
		matrix[i] = make([]int, m+1)
		dp[i] = make([]int, m+1)
	}
	dp[n+1] = make([]int, m+1)
	for i := 1; i < n+1; i++ {
		scanner.Scan()
		arrStr = strings.Split(scanner.Text(), "")
		for g, datum := range arrStr {
			if datum == "0" {
				continue
			}
			num, _ := strconv.Atoi(datum)
			matrix[i][g+1] = num
		}
	}
	for i := n; i > 0; i-- {
		for g := 1; g < m+1; g++ {
			dp[i][g] = max(dp[i+1][g], dp[i][g-1]) + matrix[i][g]
		}
	}

	s.WriteString(fmt.Sprint(dp[1][m]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

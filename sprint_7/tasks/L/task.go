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

type Heap struct {
	weight int
	cap    []int
}

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	var err error
	var arrStr []string
	var n, m int

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

	scanner.Scan()
	arrStr = strings.Fields(scanner.Text())
	previous := Heap{
		weight: 0,
		cap:    make([]int, m+1),
	}

	var goldHeap Heap
	for i := 0; i < n; i++ {
		weight, _ := strconv.Atoi(arrStr[i])
		goldHeap = Heap{
			weight: weight,
			cap:    make([]int, m+1),
		}
		for g := 0; g < m+1; g++ {
			if g-weight >= 0 {
				goldHeap.cap[g] = max(previous.cap[g], weight+previous.cap[g-weight])
			} else {
				goldHeap.cap[g] = previous.cap[g]
			}

		}
		previous = goldHeap
	}

	s.WriteString(fmt.Sprint(goldHeap.cap[m]))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

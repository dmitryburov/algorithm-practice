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

// Matrix основная матрица
type Matrix [][]int

// Peak статус вершины
type Peak uint8

const (
	_ Peak = iota
	// WHITE не обработанная вершина
	WHITE
	// GRAY обнаруженная вершина
	GRAY
	// BLACK обработанная вершина
	BLACK
)

const (
	RType = 'R'
	YES   = "YES"
	NO    = "NO"
)

func main() {
	res := strings.Builder{}
	Solution(os.Stdin, &res)

	fmt.Println(res.String())
}

// Solution решение задачи
func Solution(r io.Reader, w *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	// количество городов
	n, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		log.Fatal(err)
	}

	matrix := make(Matrix, n+1)
	visited := make([]Peak, n+1)

	// соберем матрицу
	for i := 1; i <= n; i++ {
		scanner.Scan()
		railData := scanner.Text()

		// по умолчанию вершина не посещена
		visited[i] = WHITE

		for j := 0; j < len(railData); j++ {
			target := i + j + 1
			// тип дороги R
			if railData[j] == RType {
				matrix[i] = append(matrix[i], target)
			} else {
				matrix[target] = append(matrix[target], i)
			}
		}
	}

	if matrix.checkOptimal(visited, n) {
		w.WriteString(YES)
		return
	}
	w.WriteString(NO)
}

// checkOptimal обход матрицы в поисках путей
func (m Matrix) checkOptimal(visited []Peak, n int) bool {
	for i := 1; i <= n; i++ {
		if visited[i] != WHITE {
			continue
		}

		stack := []int{i}
		for len(stack) > 0 {
			current := stack[len(stack)-1]

			if visited[current] == WHITE {
				visited[current] = GRAY

				for _, target := range m[current] {
					if visited[target] == WHITE {
						stack = append(stack, target)
					} else if visited[target] == GRAY {
						return false
					}
				}

			} else {
				visited[current] = BLACK
				stack = stack[:len(stack)-1]
			}
		}
	}

	return true
}

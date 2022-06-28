package main

/*
--- Посылка
69221420

--- Принцип работы
В случае дороги R, мы имеем свзять (N)->(N+M), а в случае 'B' - обратно (N+M)->(N).
Для решения задачи используется поиск по циклам в графе: если есть цикл длины 3,
значит существует 2 способа попасть из вершины №1 в вершину №3 - через вершину №2 по одному типу дороги,
и напрямую по другому => в таком случае карта с такими путями не будет оптимальной.

Для хранения вершин используется список смежности (матрица).
Для поиска используется алгоритм поиска в глубину DFS, т.е. во время проверки соседних нод мы встречаем серую ноду,
значит мы нашли цикл и граф дорог не является оптимальным.
Дополнительно прокомментил части в коде.

--- Временная сложность
n - количество вершин
m - количество ребер в графе
Алгоритму требуется обойти все вершины, на это требуется О(n),
а так же операция по работе со списком (вставка, удаление), на это потребуется О(m).
Cложность будет O(n + m).

--- Пространственная сложность
O(n + m), тк требуется память для матрицы и списка о посещенных вершинах.
Из прошлого опыта - исключил рекурсии из алгоритма.
Ещё есть подозрение, что используется где-то лишняя память, возможно уже мерещится =)
*/

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
	// W не обработанная вершина
	W
	// G обнаруженная вершина
	G
	// B обработанная вершина
	B
)

const (
	Y = "YES"
	N = "NO"
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
		visited[i] = W

		for j := 0; j < len(railData); j++ {
			target := i + j + 1
			// тип дороги R
			if railData[j] == 'R' {
				matrix[i] = append(matrix[i], target)
			} else {
				matrix[target] = append(matrix[target], i)
			}
		}
	}

	if checkOptimal(matrix, visited, n) {
		w.WriteString(Y)
		return
	}
	w.WriteString(N)
}

// checkOptimal обход матрицы в поисках путей
func checkOptimal(matrix Matrix, visited []Peak, n int) bool {
	for i := 1; i <= n; i++ {
		if visited[i] != W {
			continue
		}

		stack := []int{i}
		for len(stack) > 0 {
			current := stack[len(stack)-1]

			if visited[current] == W {
				visited[current] = G

				for _, target := range matrix[current] {
					if visited[target] == W {
						stack = append(stack, target)
					} else if visited[target] == G {
						return false
					}
				}

			} else {
				visited[current] = B
				stack = stack[:len(stack)-1]
			}
		}
	}

	return true
}

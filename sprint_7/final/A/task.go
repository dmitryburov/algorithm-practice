package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

// Solution решение задачи
func Solution(r io.Reader, w *strings.Builder) {
	scanner := bufio.NewScanner(r)

	// считаем строки в байты (!)
	scanner.Scan()
	s := scanner.Bytes()
	scanner.Scan()
	t := scanner.Bytes()

	lenS, lenT := len(s), len(t)

	// fix базовый случай
	if bytes.Equal(s, t) {
		w.WriteString("0")
		return
	}

	matrix := make([]int, lenS+1)
	for i := 1; i < len(matrix); i++ {
		matrix[i] = i
	}

	for i := 1; i <= lenT; i++ {
		prev := i
		for j := 1; j <= lenS; j++ {
			if t[i-1] != s[j-1] {
				matrix[j-1] = getMinimum(matrix[j]+1, getMinimum(matrix[j-1]+1, prev+1))
			}
			matrix[j-1], prev = prev, matrix[j-1]
		}
		matrix[lenS] = prev
	}

	// результат
	w.WriteString(fmt.Sprint(matrix[lenS]))
}

// getMinimum сравнение минимального числа
func getMinimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

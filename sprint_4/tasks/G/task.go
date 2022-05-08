package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int

	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	input.Scan()
	n, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	results := strings.Fields(strings.TrimSpace(input.Text()))

	fmt.Println(solution(n, results))
}

func solution(n int, results []string) int {
	var max int
	var allSum int

	resultIdx := make(map[int][]int)
	resultIdx[0] = append(resultIdx[0], 0)

	for i := 0; i < n; i++ {
		if results[i] == "0" {
			allSum++
			resultIdx[allSum] = append(resultIdx[allSum], i+1)
		} else {
			allSum--
			resultIdx[allSum] = append(resultIdx[allSum], i+1)
		}
	}
	for _, v := range resultIdx {
		if len(v) == 0 {
			continue
		}
		curr := v[len(v)-1] - v[0]
		if curr > max {
			max = curr
		}
	}

	return max
}

// getInputData подготовка входных данных
func getInputData() (scan *bufio.Scanner, err error) {
	var input *os.File
	const maxCapacity = 1024 * 1024 * 10

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}

	scanner := bufio.NewScanner(input)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	scanner.Split(bufio.ScanLines)

	return scanner, nil
}

// getInputFromFile получение ввода из файла
func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

// showError вывод ошибки
func showError(err interface{}) {
	panic(err)
}

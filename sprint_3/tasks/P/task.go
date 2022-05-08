package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		return
	}

	input.Scan()
	arrStr := strings.Split(input.Text(), " ")

	if n, err = solution(n, arrStr); err != nil {
		showError(err)
	} else {
		fmt.Println(n)
	}
}

func solution(n int, arr []string) (int, error) {
	parts, first := 0, 0
	for i := 0; i < n; i++ {

		aux := make([]int, i+1)
		for j := first; j <= i; j++ {
			nn, err := strconv.Atoi(arr[j])
			if err != nil {
				return 0, err
			}
			aux[j] = nn
		}

		sort.Ints(aux)

		isPart := true
		m := first
		for i := first; i < len(aux); i++ {
			if aux[i] != m {
				isPart = false
				break
			}
			m++
		}

		if isPart {
			parts++
			first = i + 1
		}
	}

	return parts, nil
}

// getInputData подготовка входных данных
func getInputData() (scan *bufio.Scanner, err error) {
	var input *os.File

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}

	scanner := bufio.NewScanner(input)
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

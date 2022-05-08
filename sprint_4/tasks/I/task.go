package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	var n, m int
	var nArr, mArr []string

	input.Scan()
	n, err = strconv.Atoi(strings.TrimSpace(input.Text()))
	if err != nil {
		showError(err)
	}

	input.Scan()
	nArr = make([]string, 0, n)
	nArr = strings.Fields(strings.TrimSpace(input.Text()))

	input.Scan()
	mArr = make([]string, 0, m)
	m, err = strconv.Atoi(strings.TrimSpace(input.Text()))
	if err != nil {
		showError(err)
	}

	input.Scan()
	mArr = strings.Fields(strings.TrimSpace(input.Text()))

	fmt.Println(solution(nArr, mArr))
}

func solution(nArr, mArr []string) int {
	hashMap := map[string][]int{}
	for i, v := range nArr {
		hashMap[v] = append(hashMap[v], i)
	}

	var (
		max,
		currMax int
	)

	var idx []int
	var searchIdx = -1

	actualized := func() {
		idx = nil
		if currMax > max {
			max = currMax
		}
		currMax = 0
		searchIdx = -1
	}

	for _, v := range mArr {
		if hashMap[v] != nil {
			if searchIdx == -1 && idx != nil {
				if si, ok := some(hashMap[v], idx); ok {
					searchIdx = si
					currMax++
				} else {
					actualized()
					if len(hashMap[v]) == 1 {
						searchIdx = hashMap[v][0]
						currMax++
					} else {
						idx = hashMap[v]
						currMax++
					}
				}
			} else if searchIdx == -1 {
				if len(hashMap[v]) == 1 {
					searchIdx = hashMap[v][0]
					currMax++
				} else {
					idx = hashMap[v]
					currMax++
				}
			} else {
				if r := binary(hashMap[v], searchIdx+1); r != -1 {
					searchIdx++
					currMax++
				} else {
					actualized()
				}
			}
		} else {
			actualized()
		}
	}

	if currMax > max {
		max = currMax
	}
	return max
}

func some(a, b []int) (int, bool) {
	for i, v := range a {
		if binary(b, v-1) != -1 {
			return i, true
		}
	}
	return -1, false
}

func binary(a []int, s int) int {
	logA := int(math.Log2(float64(len(a)))) + 2
	low := 0
	high := len(a)

	for i := 0; i < logA; i++ {
		middle := (low + high) >> 1
		if a[middle] == s {
			return middle
		}
		if a[middle] > s {
			high = middle
		} else {
			low = middle
		}
	}
	return -1
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	input.Scan()
	str := strings.TrimSpace(input.Text())

	fmt.Println(solution(str))
}

func solution(str string) int {
	var maxLen int

	for i, _ := range str {
		subStr := str[i:]
		chars := make(map[int32]bool)

		var j int
		for j = 0; j < len(subStr); j++ {
			if _, ok := chars[int32(subStr[j])]; ok {
				if j > maxLen {
					maxLen = j
				}
				break
			}

			chars[int32(subStr[j])] = true
		}

		if j > maxLen {
			maxLen = j
		}
	}

	return maxLen
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

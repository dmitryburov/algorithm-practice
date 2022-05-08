package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, m int
	var s string

	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	input.Scan()
	a, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	m, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	s = strings.TrimSpace(input.Text())

	fmt.Println(solution(a, m, s))

}

func solution(a, m int, s string) int {
	l := len(s)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return int(s[0]) % m
	}

	result := ((int(s[0]) * a) + int(s[1])) % m
	for i := 2; i < l; i++ {
		result = ((result * a) + int(s[i])) % m

	}
	return result
}

// getInputData подготовка входных данных
func getInputData() (scan *bufio.Scanner, err error) {
	var input *os.File
	const maxCapacity = 1024 * 1024

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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	var n, s int

	input.Scan()
	n, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	s, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	arr := strings.Fields(strings.TrimSpace(input.Text()))

	res := solution(n, s, arr)
	for i := 0; i < len(res); i++ {
		fmt.Println(strings.Join(res[i], " "))
	}
}

func solution(n, s int, arr []string) (result [][]string) {
	return
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

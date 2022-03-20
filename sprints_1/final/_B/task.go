package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	inp, err := getInputFromFile()
	if err != nil {
		showError(err)
	}

	n, fields, err := parseInputData(bufio.NewReader(inp))
	if err != nil {
		showError(err)
	}

	result := checkResult(n, fields)
	fmt.Println(result)
}

func checkResult(n int, fields []int) int {

	return 0
}

// parseInputData парсинг данных из input
func parseInputData(reader *bufio.Reader) (n int, fields []int, err error) {

	return
}

// getInputFromFile получение input из файла
func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

// showError вывод ошибки
func showError(err error) {
	panic(err)
}

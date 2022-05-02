package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int

	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	// обрабатываем документы
	input.Scan()
	n, err = strconv.Atoi(input.Text())
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		input.Scan()
		fmt.Println(input.Text())
	}

	// обрабатываем запрос
	input.Scan()
	n, err = strconv.Atoi(input.Text())
	if err != nil {
		return
	}
	for i := 0; i < n; i++ {
		input.Scan()
		fmt.Println(input.Text())
	}

}

func testMain(docs, query []string) (result [][]int) {
	return result
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type input struct {
	data,
	search []string
}

func main() {
	data, err := getInputData()
	if err != nil {
		showError(err)
	}

	fmt.Println(data)
}

// getInputData подготовка входных данных
func getInputData() (data *input, err error) {
	var input *os.File
	var bufStr string
	var n int

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	bufStr = scanner.Text()
	n, err = strconv.Atoi(bufStr)
	if err != nil {
		return
	}

	fmt.Println(n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		bufStr = scanner.Text()
		fmt.Println(bufStr)
	}

	scanner.Scan()
	bufStr = scanner.Text()
	n, err = strconv.Atoi(bufStr)
	if err != nil {
		return
	}

	fmt.Println(n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		bufStr = scanner.Text()
		fmt.Println(bufStr)
	}

	return
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

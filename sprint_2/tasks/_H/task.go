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

	result := solution(input)
	fmt.Println(result)
}

// solution решение задачи
func solution(input string) (result string) {
	for strings.Contains(input, "[]") ||
		strings.Contains(input, "{}") ||
		strings.Contains(input, "()") {
		input = strings.ReplaceAll(input, "[]", "")
		input = strings.ReplaceAll(input, "{}", "")
		input = strings.ReplaceAll(input, "()", "")
	}

	if len(input) > 0 {
		return "False"
	}

	return "True"
}

// getInputData парсинг входных данных
func getInputData() (input string, err error) {

	inp, err := getInputFromFile()
	if err != nil {
		return
	}
	defer func(inp *os.File) {
		_ = inp.Close()
	}(inp)

	reader := bufio.NewReader(inp)

	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	defer reader.Reset(reader)
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

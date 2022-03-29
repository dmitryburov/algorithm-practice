package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	n, numbers, err := getInputData()
	if err != nil {
		showError(err)
	}

	solution(n, numbers)

	if err = outputFile(strings.Join(numbers, " ")); err != nil {
		showError(err)
	}
}

// solution решение задачи
func solution(n int, numbers []string) []string {
	var (
		zeroIndex = make(map[int]int)
		counter   = n * 2
	)

	// обход слева
	for i := 0; i < n; i++ {
		if numbers[i] == "0" {
			if i+1 < n && numbers[i+1] == "0" {
				continue
			}
			zeroIndex[i] = counter
			counter = 1
			continue
		}

		if counter < n {
			numbers[i] = fmt.Sprint(counter)
		} else {
			numbers[i] = fmt.Sprint(n)
		}

		counter++
	}

	// обход справа
	var contain int
	counter = 999999
	for i := n - 1; i >= 0; i-- {
		if numbers[i] == "0" {
			if i+1 < n && numbers[i+1] == "0" {
				continue
			}
			counter = 1
			contain = zeroIndex[i]
			continue
		}

		if contain > 0 {
			if counter < (contain - counter) {
				numbers[i] = fmt.Sprint(counter)
			}
			counter++
		}
	}

	return numbers
}

// getInputData парсинг входных данных
func getInputData() (n int, numbers []string, err error) {

	input, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	reader := bufio.NewReader(input)

	strNum, _, _ := reader.ReadLine()
	n, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	numbers = make([]string, n)
	strNums, _ := reader.ReadString('\n')
	numbers = strings.Split(strings.TrimSpace(strNums), " ")

	// clear bufio
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

// outputFile вывод output в файл
func outputFile(data string) error {
	f, err := os.Create("output.txt")
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		_ = f.Close()
	}(f)

	_, err = f.WriteString(data + "\n")
	if err != nil {
		return err
	}

	return nil
}

// showError обработка ошибки
func showError(err interface{}) {
	panic(err)
}

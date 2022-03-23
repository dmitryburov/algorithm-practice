package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ID посылки:
// Предыдущий ID посылки: 66269683
func main() {

	n, numbers, err := getInputData()
	if err != nil {
		showError(err)
	}

	solution(n, numbers)

	if err = outputFile(strings.Join(numbers, "")); err != nil {
		showError(err)
	}
}

// solution решение задачи
func solution(n int, numbers []string) []string {
	var (
		zeroIndex = make(map[int]int)
		findZero  bool
		lastZero  = 1
	)

	// обход слева
	for i := 0; i < n; i++ {
		if numbers[i] == "0" {
			zeroIndex[i] = 0
			findZero = true
			lastZero = 1
			continue
		}

		if findZero {
			numbers[i] = fmt.Sprintf("%d", lastZero)
			lastZero++
		} else {
			numbers[i] = fmt.Sprintf("%d", n)
		}
	}

	// обход справа
	lastZero = 1
	for i := n - 1; i >= 0; i-- {
		if numbers[i] == "0" {
			lastZero = 1
			delete(zeroIndex, len(zeroIndex)-1) // remove last
			continue
		}

		if len(zeroIndex) > 0 {
			if i - zeroIndex[] > lastZero {
				numbers[i] = fmt.Sprintf("%d", lastZero)
			}
		}

		lastZero++

		if findLastZero {
			if numbers[i] > lastZero {
				numbers[i] = fmt.Sprintf("%d", lastZero)
			}
			lastZero++
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
func showError(err error) {
	// не буду сильно заморачиваться, но можно облагородить =)
	panic(err)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int

	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	input.Scan()
	n, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	arrStr := strings.Fields(input.Text())
	fmt.Println(len(arrStr))
	if _, err := solution(n, arrStr); err != nil {
		showError(err)
	} else {
		//fmt.Println(strings.Join(res, " "))
	}
}

func solution(n int, arr []string) ([]string, error) {
	if n == 0 {
		return nil, nil
	}
	counter := make([]int, 0, 3)

	for i := 0; i < len(arr); i++ {
		if arr[i] == "" {
			continue
		}

		vInt, err := strconv.Atoi(arr[i])
		if err != nil {
			return nil, err
		}
		counter[vInt] += 1
	}

	result := make([]string, n)
	fmt.Println(counter)
	i := 0
	for g, v := range counter {
		for v > 0 {
			result[i] = strconv.Itoa(g)
			v--
			i++
		}
	}

	return result, nil
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

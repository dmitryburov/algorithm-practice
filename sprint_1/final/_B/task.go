package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ID посылки: 66293083
func main() {

	const matrixCnt = 4

	inp, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	defer inp.Close()

	reader := bufio.NewReader(inp)

	strNum, _, _ := reader.ReadLine()
	k, err := strconv.Atoi(string(strNum))
	if err != nil {
		showError(err)
	}

	// к могут нажать каждый из игроков
	k *= 2

	var numbers [][]string
	for i := 0; i < matrixCnt; i++ {
		strNums, _ := reader.ReadString('\n')
		numbers = append(numbers, strings.Split(strings.TrimSpace(strNums), ""))
	}
	defer reader.Reset(reader)

	var result int
	var matrix = make(map[string]int, 9)

	// собираем совпадения в матрице
	for i := 0; i < matrixCnt; i++ {
		for i2 := 0; i2 < matrixCnt; i2++ {
			if numbers[i][i2] != "." {
				if _, ok := matrix[numbers[i][i2]]; !ok {
					matrix[numbers[i][i2]] = 1
				} else {
					matrix[numbers[i][i2]]++
				}
			}
		}
	}

	// проверям совпадения по времени t
	for i := range matrix {
		if k >= matrix[i] {
			result++
		}
	}

	fmt.Println(result)
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

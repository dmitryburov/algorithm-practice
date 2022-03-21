package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

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
	for i := 0; i < 4; i++ {
		strNums, _ := reader.ReadString('\n')
		numbers = append(numbers, strings.Split(strings.TrimSpace(strNums), ""))
	}
	defer reader.Reset(reader)

	var result int
	var matrix = make(map[string]int, 10)

	// собираем совпадения в матрице
	for i := 0; i < 4; i++ {
		if numbers[i][i] != "." {

		}
		if _, ok := matrix[numbers]; !ok {

		}
	}

	// проверям совпадения по времени t

	fmt.Println(matrix)
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

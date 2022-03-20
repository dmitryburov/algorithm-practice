package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	defer input.Close()

	n, numbers, err := parseDataInput(bufio.NewReader(input))
	if err != nil {
		showError(err)
	}

	result := resultTask(n, numbers)
	if err = outputFile(result); err != nil {
		showError(err)
	}
}

// resultTask поиск расстояния до соседей
func resultTask(n int, numbers []string) string {

	var needCalc bool
	var zeroCount = 1

	for i := 0; i < n; i++ {
		if numbers[i] == "0" {
			if needCalc {
				needCalc = false
				if zeroCount%2 == 1 {
					zeroCount = (zeroCount + 1) / 2
				} else {
					zeroCount = zeroCount / 2
				}
				i -= zeroCount
				continue
			} else {
				for i2 := i; i2 > 0; i2-- {
					numbers[i2] =

				}
			}

			needCalc = true
			zeroCount = 1
		} else {
			numbers[i] = fmt.Sprintf("%d", zeroCount)
			zeroCount++
		}
	}

	return strings.Join(numbers, " ")
}

// parseDataInput парсинг параметров согласно задаче
func parseDataInput(reader *bufio.Reader) (n int, numbers []string, err error) {

	defer reader.Reset(reader)

	strNum, _, _ := reader.ReadLine()
	n, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	strNums, _ := reader.ReadString('\n')
	numbers = strings.Split(strings.TrimSpace(strNums), " ")

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

	defer f.Close()

	_, err = f.WriteString(data + "\n")
	_, err = f.WriteString("0 0 1 0 0 0 0 1 0 0 1 2 1 0 1 0 1 0 0 0 # etalon" + "\n")
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

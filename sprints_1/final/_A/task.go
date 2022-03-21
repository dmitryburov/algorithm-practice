package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ID посылки: 66269683
// Комментарий:
//    Cделал алогиритм линейным, думал-думал как же го реализовать по сложности O(log(n),
//    тк при нагрузке время и память расходовалась прилично.
// 	  Стал переживать что не успею =)
func main() {

	input, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	// fix close file
	defer input.Close()

	reader := bufio.NewReader(input)

	strNum, _, _ := reader.ReadLine()
	n, err := strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	numbers := make([]string, n)
	strNums, _ := reader.ReadString('\n')
	numbers = strings.Split(strings.TrimSpace(strNums), " ")

	// fix close bufio
	defer reader.Reset(reader)

	var (
		findLastZero bool
		lastZero     = 1
	)

	// обход слева
	for i := 0; i < n; i++ {
		if numbers[i] == "0" {
			lastZero = 1
			findLastZero = true
			continue
		}

		if findLastZero {
			numbers[i] = fmt.Sprintf("%d", lastZero)
			lastZero++
		} else {
			numbers[i] = fmt.Sprintf("%d", n)
		}
	}

	findLastZero = false
	lastZero = 1

	// обход справа
	for i := n - 1; i >= 0; i-- {
		if numbers[i] == "0" {
			lastZero = 1
			findLastZero = true
			continue
		}

		if findLastZero {
			currNum, _ := strconv.Atoi(numbers[i])
			if currNum > lastZero {
				numbers[i] = fmt.Sprintf("%d", lastZero)
			}

			lastZero++
		}
	}

	if err = outputFile(strings.Join(numbers, " ")); err != nil {
		showError(err)
	}
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

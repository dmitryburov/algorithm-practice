package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	responseYes = "YES"
	responseNo  = "NO"
)

func main() {
	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	var s, t string

	input.Scan()
	s = strings.TrimSpace(input.Text())

	input.Scan()
	t = strings.TrimSpace(input.Text())

	fmt.Println(solution(s, t))
}

func solution(s, t string) string {
	if len(s) != len(t) {
		return responseNo
	}

	sToT := make(map[string]string)
	tToS := make(map[string]string)

	for i := 0; i < len(s); i++ {
		sL := string(s[i])
		tL := string(t[i])
		if sToT[sL] == "" && tToS[tL] == "" {
			sToT[sL] = tL
			tToS[tL] = sL
		}
		if sToT[sL] != tL || tToS[tL] != sL {
			return responseNo
		}
	}

	return responseYes
}

// getInputData подготовка входных данных
func getInputData() (scan *bufio.Scanner, err error) {
	var input *os.File
	const maxCapacity = 1024 * 1024 * 10

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}

	scanner := bufio.NewScanner(input)

	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

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

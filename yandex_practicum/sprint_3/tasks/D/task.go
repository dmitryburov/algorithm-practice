package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	var nArr, mArr []string

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
	nArr = strings.Fields(input.Text())

	input.Scan()
	m, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	mArr = strings.Fields(input.Text())

	if res, err := solution(n, m, nArr, mArr); err != nil {
		showError(err)
	} else {
		fmt.Println(res)
	}
}

func solution(n, m int, nArr, mArr []string) (int, error) {
	sortGreed := merge(nArr)
	sortCookie := merge(mArr)
	cookieIndex := 0
	count := 0

	for _, v := range sortGreed {
		if cookieIndex >= len(sortCookie) {
			break
		}
		for ; cookieIndex < len(sortCookie); cookieIndex++ {
			if sortCookie[cookieIndex] >= v {
				count++
				cookieIndex++
				break
			}
		}
	}

	return count, nil
}

func merge(s []string) []int {

	if len(s) == 1 {
		i, _ := strconv.Atoi(s[0])
		result := []int{i}
		return result
	}
	middle := len(s) / 2
	a := merge(s[:middle])
	b := merge(s[middle:])

	aIndex := 0
	bIndex := 0
	result := make([]int, 0, len(a)+len(b))
	for aIndex < len(a) || bIndex < len(b) {
		if aIndex == len(a) {
			result = append(result, b[bIndex])
			bIndex++
			continue
		}
		if bIndex == len(b) {
			result = append(result, a[aIndex])
			aIndex++
			continue
		}
		if a[aIndex] <= b[bIndex] {
			result = append(result, a[aIndex])
			aIndex++
		} else {
			result = append(result, b[bIndex])
			bIndex++
		}
	}
	return result
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

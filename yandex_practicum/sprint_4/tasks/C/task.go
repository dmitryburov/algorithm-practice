package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, m, t int
	var s string
	var arr []string

	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	input.Scan()
	a, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	m, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	input.Scan()
	s = input.Text()

	input.Scan()
	t, err = strconv.Atoi(input.Text())
	if err != nil {
		showError(err)
	}

	for i := 0; i < t; i++ {
		input.Scan()
		arr = append(arr, strings.TrimSpace(input.Text()))
	}

	if res, err := solution(a, m, t, s, arr); err != nil {
		showError(err)
	} else {
		for i := 0; i < len(res); i++ {
			fmt.Println(res[i])
		}
	}
}

func solution(a, m, t int, s string, q []string) ([]int, error) {

	if len(s) == 0 {
		return []int{0}, nil
	}

	degree := degreeByModule(a, m, len(s))
	pref := prefixHash(s, a, m)
	result := make([]int, t)

	for i := 0; i < t; i++ {
		arrStr := strings.Fields(q[i])
		aI, _ := strconv.Atoi(arrStr[0])
		bI, _ := strconv.Atoi(arrStr[1])

		if aI-1 == 0 {
			result[i] = pref[bI]
		} else {
			v := pref[bI] - pref[aI-1]*degree[bI-aI+1]
			if v < 0 {
				v %= m
				v += m
				v %= m
			}
			result[i] = v
		}
	}

	return result, nil
}

func degreeByModule(val, mod, n int) []int {
	degrees := make([]int, n+1)
	degrees[0] = 1

	for i := 1; i <= n; i++ {
		res := degrees[i-1] * val
		degrees[i] = res % mod
	}

	return degrees
}

func prefixHash(s string, a, m int) []int {
	pref := make([]int, len(s)+1)
	pref[0] = 0

	for i := 1; i <= len(s); i++ {
		pref[i] = (pref[i-1]*a + int(s[i-1])) % m
	}

	return pref
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

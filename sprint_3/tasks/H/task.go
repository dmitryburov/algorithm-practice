package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	arr, err := getInputData()
	if err != nil {
		showError(err)
	}

	fmt.Println(getBigNum(arr))
}

func getBigNum(arr []string) string {

	var n = len(arr)
	var changed bool

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if less(arr[j], arr[j+1]) {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				changed = true
			}
		}

		if !changed {
			break
		}

	}

	return strings.Join(arr, "")
}

func less(a, b string) bool {
	c, err := strconv.Atoi(a + b)
	if err != nil {
		showError(err)
	}
	d, err := strconv.Atoi(b + a)
	if err != nil {
		showError(err)
	}

	return c < d
}

func getInputData() (arr []string, err error) {
	var input *os.File
	var btStr []byte

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	var n int
	reader := bufio.NewReader(input)
	btStr, _, err = reader.ReadLine()
	n, err = strconv.Atoi(string(btStr))
	if err != nil {
		return
	}

	arr = make([]string, n)

	btStr, _, err = reader.ReadLine()
	arr = strings.Split(string(btStr), " ")

	// clear bufio
	defer reader.Reset(reader)
	return
}

func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func showError(err interface{}) {
	panic(err)
}

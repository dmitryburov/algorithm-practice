package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	items, err := getInputData()
	if err != nil {
		showError(err)
	}

	for i := range items {
		fmt.Println(i, []byte(i))
	}
}

func getInputData() (items map[string]bool, err error) {
	var input *os.File
	var bufStr []byte
	var n int

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	bufStr = scanner.Text()
	n, err = strconv.Atoi(string(bufStr))
	if err != nil {
		return
	}

	items = make(map[string]bool, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		bufStr = scanner.Bytes()

		if _, ok := items[bufStr]; !ok {
			items[bufStr] = true
		}
	}

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

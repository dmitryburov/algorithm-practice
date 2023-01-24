package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input *os.File
	var bufStr string
	var err error
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
	n, err = strconv.Atoi(bufStr)
	if err != nil {
		return
	}

	items := make(map[string]bool, n)

	for i := 0; i < n; i++ {
		scanner.Scan()
		bufStr = scanner.Text()
		if _, ok := items[bufStr]; !ok {
			fmt.Println(bufStr)
			items[bufStr] = true
		}
	}
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

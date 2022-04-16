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

	bubbleSort(arr)
}

func bubbleSort(arr []int) {

	var n = len(arr)
	var changed bool

	for i := 0; i < n; i++ {
		changed = false

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				changed = true
			}
		}

		if changed ||
			!changed && i < 1 { // fix
			printArr(arr)
		}

		if !changed {
			break
		}

	}

}

func printArr(arr []int) {
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}

func getInputData() (arr []int, err error) {
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

	arr = make([]int, n)

	btStr, _, err = reader.ReadLine()
	strArr := strings.Split(string(btStr), " ")
	for i := 0; i < n; i++ {
		arr[i], _ = strconv.Atoi(strArr[i])
	}

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

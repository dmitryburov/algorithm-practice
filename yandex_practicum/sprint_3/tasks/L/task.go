package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, price, days, err := getInputData()
	if err != nil {
		showError(err)
	}

	fmt.Println(strings.Trim(fmt.Sprint(solution(n, price, days)), "[]"))
}

func solution(n, price int, days []int) []int {
	var day1, day2 int

	day1 = binarySearch(price, days, 0, n)
	day2 = binarySearch(price*2, days, 0, n)

	return []int{day1, day2}
}

func binarySearch(x int, arr []int, l, r int) int {
	if l >= r {
		return -1
	}

	if arr[l] >= x {
		return l + 1
	}

	m := l + (r-l)/2

	if arr[m] >= x {
		if m == l+1 {
			return m + 1
		} else {
			return binarySearch(x, arr, l, m+1)
		}
	} else if arr[m] > x {
		return binarySearch(x, arr, l, m)
	} else {
		return binarySearch(x, arr, m+1, r)
	}
}

func getInputData() (n, price int, daysPrice []int, err error) {

	input, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	reader := bufio.NewReader(input)

	strNum, _, _ := reader.ReadLine()
	n, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	var num int
	daysPrice = make([]int, n)
	strNums, _ := reader.ReadString('\n')
	strNumsArr := strings.Split(strings.TrimSpace(strNums), " ")
	for i := 0; i < n; i++ {
		num, _ = strconv.Atoi(strNumsArr[i])
		daysPrice[i] = num
	}

	strNum, _, _ = reader.ReadLine()
	price, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
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

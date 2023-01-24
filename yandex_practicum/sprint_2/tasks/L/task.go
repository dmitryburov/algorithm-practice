package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, k, err := getInputData()
	if err != nil {
		showError(err)
	}

	result := solution(n, k)
	fmt.Println(result)
}

func solution(n, k int) int {
	if n < 2 {
		return 1
	}

	var result int
	ab := []int{1, 1}
	k = int(math.Pow(10, float64(k)))

	for i := 0; i < n-1; i++ {
		s := (ab[0] + ab[1]) % k
		ab[0] = ab[1]
		ab[1] = s
		result = ab[1]
	}

	return result
}

// getInputData парсинг входных данных
func getInputData() (n, k int, err error) {

	reader := bufio.NewReader(os.Stdin)

	strNum, _, _ := reader.ReadLine()
	strExp := strings.Split(string(strNum), " ")

	n, err = strconv.Atoi(strExp[0])
	if err != nil {
		return
	}

	k, err = strconv.Atoi(strExp[1])
	if err != nil {
		return
	}

	defer reader.Reset(reader)
	return
}

// showError вывод ошибки
func showError(err error) {
	panic(err)
}

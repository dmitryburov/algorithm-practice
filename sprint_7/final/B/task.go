package main

/*
--- Посылка
69267329

--- Принцип работы
Дополнительные комментарии в коде.

--- Временная сложность

--- Пространственная сложность

*/
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	SUCCESS = "True"
	FAIL    = "False"
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	var (
		n       int
		err     error
		scanner = bufio.NewScanner(r)
	)

	// количество выигранных партий
	scanner.Scan()
	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	// соберем слайс заработанных в партиях очков
	data := make([]int, n)
	for scanner.Scan() {
		arrString := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(arrString); i++ {
			data[i], _ = strconv.Atoi(arrString[i])
		}
	}

	if !CheckSum(n, data) {
		s.WriteString(FAIL)
		return
	}

	s.WriteString(SUCCESS)
}

// CheckSum проверка суммы
func CheckSum(n int, data []int) bool {
	sum := GetSum(data)

	// базовый случай
	if sum%2 != 0 {
		return false
	}

	dpPrev := make([]int, sum/2+1)
	dp := make([]int, sum/2+1)
	for i := 1; i < n+1; i++ {
		for g := 1; g < sum/2+1; g++ {
			dp[g] = dpPrev[g]
			if g == data[i-1] {
				dp[g]++
			}
			if g > data[i-1] && dpPrev[g-data[i-1]] > 0 {
				dp[g]++
			}
		}
		dpPrev = dp
		dp = make([]int, sum/2+1)
	}

	return dpPrev[sum/2] > 1
}

// GetSum суммирует значения очков
func GetSum(arr []int) (sum int) {
	for i := range arr {
		sum += arr[i]
	}

	return
}

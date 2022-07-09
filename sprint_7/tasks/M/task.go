package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type BackPackItem struct {
	value  int
	weight int
}

type BackPack []BackPackItem

type Matrix [][]int

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	var err error
	var arrStr []string
	var n, m int

	scanner.Scan()
	arrStr = strings.Fields(scanner.Text())

	n, err = strconv.Atoi(arrStr[0])
	if err != nil {
		log.Fatal(err)
	}

	m, err = strconv.Atoi(arrStr[1])
	if err != nil {
		log.Fatal(err)
	}

	backPack := make(BackPack, 0, n)

	var weight, value int
	for i := 0; i < n; i++ {
		scanner.Scan()
		arrStr = strings.Fields(scanner.Text())

		weight, err = strconv.Atoi(arrStr[0])
		if err != nil {
			log.Fatal(err)
		}

		value, err = strconv.Atoi(arrStr[1])
		if err != nil {
			log.Fatal(err)
		}

		backPack = append(backPack, BackPackItem{
			value:  value,
			weight: weight,
		})
	}

	dp := make(Matrix, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		if i == 0 {
			continue
		}

		for g := 1; g < m+1; g++ {
			if g < backPack[i-1].weight {
				dp[i][g] = max(dp[i-1][g], dp[i][g-1])
			}
			if g >= backPack[i-1].weight {
				dp[i][g] = max(dp[i-1][g-backPack[i-1].weight]+backPack[i-1].value, dp[i-1][g])
			}
		}
	}

	current := dp[n][m]
	i, g := n, m
	result := make([]string, 0, 16)

	for current != 0 {
		if dp[i-1][g] == current {
			i--
			continue
		}
		if dp[i][g-1] == current {
			g--
			continue
		}

		result = append(result, strconv.Itoa(i))
		g -= backPack[i-1].weight
		i--
		current = dp[i][g]
	}

	scanner.Scan()
	s.WriteString(fmt.Sprint(len(result)) + "\n")
	s.WriteString(strings.Join(result, " "))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

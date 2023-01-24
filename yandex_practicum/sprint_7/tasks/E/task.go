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

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	x, k, nominals := getData(bufio.NewScanner(r))

	total := -1
	prev := make([]int, x+1)
	result := make([]int, x+1)

	for i := 0; i < k; i++ {
		for j := 0; j < x+1; j++ {
			result[j] = min(result[j], prev[j])
			if j-nominals[i] >= 0 && (result[j-nominals[i]] > 0 || j%nominals[i] == 0) {
				result[j] = min(result[j], result[j-nominals[i]]+1)
			}
		}

		prev = result
		result = make([]int, x+1)
	}

	if prev[x] > 0 {
		total = prev[x]
	}

	s.WriteString(fmt.Sprint(total))
}

func min(a, b int) int {
	if a == 0 || b == 0 {
		if a > b {
			return a
		}
	} else if a < b {
		return a
	}

	return b
}

func getData(scanner *bufio.Scanner) (x, k int, nominal []int) {
	var err error

	scanner.Scan()
	x, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	k, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	strFields := strings.Fields(scanner.Text())
	nominal = make([]int, k)

	for i := 0; i < k; i++ {
		nominal[i], err = strconv.Atoi(strFields[i])
		if err != nil {
			log.Fatal(err)
		}
	}

	return
}

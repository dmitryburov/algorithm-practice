package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(strings.TrimSpace(s.String()))
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	const bufCapacity = 100000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	scanner.Scan()
	str := scanner.Text()

	dp := make([]int, len(str))
	s.WriteString("0 ")
	for i := 1; i < len(str); i++ {
		k := dp[i-1]
		for k > 0 && str[i] != str[k] {
			k = dp[k-1]
		}
		if str[i] == str[k] {
			dp[i] = k + 1
			s.WriteString(strconv.Itoa(dp[i]) + " ")
		} else {
			dp[i] = 0
			s.WriteString("0 ")
		}
	}
}

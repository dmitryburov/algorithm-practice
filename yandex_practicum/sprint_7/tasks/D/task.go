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

const MOD = 1000000007

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)
	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i < n+1; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % MOD
	}

	s.WriteString(fmt.Sprint(dp[n]))
}

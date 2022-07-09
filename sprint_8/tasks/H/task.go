package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, w *strings.Builder) {
	scanner := bufio.NewScanner(r)

	const bufCapacity = 10000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	var text, s, t string

	scanner.Scan()
	text = scanner.Text()
	scanner.Scan()
	s = scanner.Text()
	scanner.Scan()
	t = scanner.Text()

	sentinel := s + "#" + text
	patternLen := len(s)
	prefixLen := patternLen + 1

	dp := make([]int, len(sentinel))
	result := make([]bool, len(text))
	for i := prefixLen; i < len(sentinel); i++ {
		k := dp[i-1]
		for sentinel[i] != sentinel[k] && k > 0 {
			k = dp[k-1]
		}
		if sentinel[k] != sentinel[i] {
			dp[i] = k
		} else {
			dp[i] = k + 1
			if dp[i] == patternLen {
				result[i-patternLen*2] = true
			}
		}

	}

	for i := 0; i < len(text); i++ {
		if result[i] {
			w.WriteString(t)
			i += patternLen - 1
		} else {
			w.WriteByte(text[i])
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
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
	aBuilder := strings.Builder{}
	for _, symbol := range scanner.Text() {
		if unicode.Is(unicode.White_Space, symbol) {
			break
		}
		if symbol%2 == 0 {
			aBuilder.WriteRune(symbol)
		}
	}

	scanner.Scan()
	bBuilder := strings.Builder{}
	for _, symbol := range scanner.Text() {
		if unicode.Is(unicode.White_Space, symbol) {
			break
		}
		if symbol%2 == 0 {
			bBuilder.WriteRune(symbol)
		}
	}

	s.WriteString(compare(aBuilder.String(), bBuilder.String()))
}

func compare(a, b string) string {
	if a == b {
		return "0"
	}
	for i := 0; i < len(a); i++ {
		if i == len(b) {
			return "1"
		}
		if a[i] == b[i] {
			continue
		}
		if a[i] < b[i] {
			return "-1"
		} else {
			return "1"
		}
	}
	return "-1"
}

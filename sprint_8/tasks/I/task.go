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

	fmt.Println(strings.TrimSpace(s.String()))
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	const bufCapacity = 100000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	scanner.Scan()
	str := scanner.Text()

	prefix := string(str[0])
	prefixLength := len(prefix)
	i := 1

	for i+prefixLength <= len(str) {
		if prefix == str[i:i+prefixLength] {
			i += prefixLength
		} else {
			i++
			prefix = str[:i]
			prefixLength = len(prefix)
		}
	}
	if len(prefix) > len(str)/2 {
		prefix = str
		prefixLength = len(prefix)
	}

	s.WriteString(fmt.Sprint(len(str) / prefixLength))
}

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

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	fields := strings.Fields(scanner.Text())
	for i, g := 0, len(fields)-1; i < g; i, g = i+1, g-1 {
		fields[i], fields[g] = fields[g], fields[i]
	}

	s.WriteString(strings.Join(fields, " "))
}

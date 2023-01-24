package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)

	var find string
	var values [][]string

	for i := 1; i <= 2; i++ {
		scanner.Scan()

		parse := strings.Split(scanner.Text(), "")
		sort.StringSlice(parse).Sort()

		values = append(values, parse)
	}

	for i := 0; i < len(values[1]); i++ {
		if i+1 > len(values[0]) {
			find = values[1][i]
			break
		}
		if values[1][i] != values[0][i] {
			find = values[1][i]
			break
		}
	}

	_, _ = writer.WriteString(find)
	_, _ = writer.WriteString("\n")
	_ = writer.Flush()
}

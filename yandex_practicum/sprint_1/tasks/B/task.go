package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	line := scanner.Text()

	var result = "WIN"
	var values = strings.Split(line, " ")
	var last int

	for i := 0; i < len(values); i++ {
		num, _ := strconv.Atoi(values[i])
		if num%2 == 0 {
			if last == 1 {
				result = "FAIL"
				break
			}
			last = 2
		} else {
			if last == 2 {
				result = "FAIL"
				break
			}
			last = 1
		}
	}

	_, _ = writer.WriteString(result)
	_, _ = writer.WriteString("\n")
	_ = writer.Flush()
}

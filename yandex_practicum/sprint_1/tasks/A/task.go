package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var a, b, c, x int

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	line := scanner.Text()

	values := strings.Split(line, " ")

	a, _ = strconv.Atoi(values[0])
	x, _ = strconv.Atoi(values[1])
	b, _ = strconv.Atoi(values[2])
	c, _ = strconv.Atoi(values[3])

	result := a*(x*x) + b*x + c

	_, _ = writer.WriteString(strconv.Itoa(result))
	_, _ = writer.WriteString("\n")
	_ = writer.Flush()
}

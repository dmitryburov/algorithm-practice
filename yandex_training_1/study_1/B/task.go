package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	a, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	b, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	c, _ := strconv.Atoi(scanner.Text())

	fmt.Println(solution(a, b, c))
}

func solution(a, b, c int) string {
	if a+b <= c || b+c <= a || a+c <= b {
		return "NO"
	}

	return "YES"
}

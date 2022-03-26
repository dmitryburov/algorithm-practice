package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scan(&n)
	fmt.Println(solution(n))
}

// solution решение задачи
func solution(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	a, b := solution(n-1), solution(n-2)
	return a + b
}

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
	if n <= 1 {
		return 1
	}

	return solution(n-1) + solution(n-2)
}

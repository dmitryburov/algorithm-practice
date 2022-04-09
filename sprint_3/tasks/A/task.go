package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)

	solution(n, "", 0, 0)
}

func solution(n int, line string, l, r int) {
	if l == n && r == n {
		fmt.Println(line)
	} else {
		if l < n {
			solution(n, line+"(", l+1, r)
		}
		if r < l {
			solution(n, line+")", l, r+1)
		}
	}
}

package main

import (
	"fmt"
	"strings"
)

func main() {

	var n int
	fmt.Scan(&n)

	res := strings.Builder{}
	solution(n, "", 0, 0, &res)
	fmt.Println(strings.TrimRight(res.String(), "\n"))
}

func solution(n int, line string, l, r int, s *strings.Builder) {
	if l == n && r == n {
		s.WriteString(line + "\n")
	} else {
		if l < n {
			solution(n, line+"(", l+1, r, s)
		}
		if r < l {
			solution(n, line+")", l, r+1, s)
		}
	}
}

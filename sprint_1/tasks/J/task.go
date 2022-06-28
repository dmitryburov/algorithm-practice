package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var res []string
	var i = 2

	for i*i <= n {
		if n%i == 0 {
			n /= i
			res = append(res, fmt.Sprintf("%d", i))
		} else {
			i += 1
		}
	}

	if n > 1 {
		res = append(res, fmt.Sprintf("%d", n))
	}

	fmt.Println(strings.Join(res, " "))
}

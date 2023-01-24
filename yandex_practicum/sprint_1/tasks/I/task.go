package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scan(&n)

	if res := numToDegree(n, 4); res {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func numToDegree(num int, degree int) bool {
	for num%degree == 0 {
		num /= degree
	}

	return num == 1
}

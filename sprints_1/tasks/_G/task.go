package main

import (
	"fmt"
)

func main() {

	var n int

	fmt.Scan(&n)
	fmt.Println(numToBinary(n))
}

func numToBinary(num int) string {
	var binaryStr string
	var n = num

	for n > 0 {
		binaryStr = fmt.Sprintf("%d%s", n%2, binaryStr)
		n = n / 2
	}

	return binaryStr
}

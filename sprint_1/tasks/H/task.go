package main

import (
	"fmt"
)

func main() {

	var one, two int

	fmt.Scanf("%d", &one)
	fmt.Scanf("%d", &two)

	fmt.Println("One:", one)
	fmt.Println("Two:", two)

}

func add(a, b int) int {
	if b == 0 {
		return a
	}
	sum := a ^ b
	carry := (a & b) << 1
	return add(sum, carry)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	a, b := getInputData(bufio.NewScanner(os.Stdin))
	fmt.Println(solution(a, b))
}

func solution(a, b int) int {
	return a + b
}

func getInputData(scanner *bufio.Scanner) (a, b int) {
	const bufCapacity = 10000000 // fix long
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	var err error

	scanner.Scan()
	a, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	b, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	return
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	isAuto   = "auto"
	isFreeze = "freeze"
	isHeat   = "heat"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	temperatures := strings.Fields(scanner.Text())
	t1, _ := strconv.Atoi(temperatures[0])
	t2, _ := strconv.Atoi(temperatures[1])

	scanner.Scan()
	action := scanner.Text()

	fmt.Println(solution(t1, t2, action))
}

func solution(t1, t2 int, action string) int {
	switch {
	case action == isAuto,
		action == isFreeze && t1 > t2,
		action == isHeat && t1 < t2:
		return t2
	}

	return t1
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const MOD = 1000000007

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	var err error
	var n, k int

	scanner.Scan()
	arrStr := strings.Fields(scanner.Text())

	n, err = strconv.Atoi(arrStr[0])
	if err != nil {
		log.Fatal(err)
	}

	k, err = strconv.Atoi(arrStr[1])
	if err != nil {
		log.Fatal(err)
	}

	ladder := make([]int, n+1)
	ladder[0] = 0
	ladder[1] = 1

	for i := 2; i < len(ladder); i++ {
		if k >= i {
			for j := 1; j < i; j++ {
				ladder[i] += ladder[j]
			}
		} else {
			for j := i - k; j < i; j++ {
				ladder[i] += ladder[j]
			}
		}

		ladder[i] = ladder[i] % MOD
	}

	s.WriteString(fmt.Sprint(ladder[n]))
}

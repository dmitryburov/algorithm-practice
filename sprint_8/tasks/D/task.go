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

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	const bufCapacity = 10000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	var err error
	var n int

	scanner.Scan()
	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	var last, str string
	for i := 0; i < n; i++ {
		scanner.Scan()
		str = scanner.Text()

		if i == 0 {
			last = str
			continue
		}

		var j, g int
		for j < len(last) && g < len(str) && last[j] == str[g] {
			j++
			g++
		}
		last = str[:g]
	}

	s.WriteString(fmt.Sprint(len(last)))
}

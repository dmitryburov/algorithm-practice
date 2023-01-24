package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(strings.TrimSpace(s.String()))
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	const bufCapacity = 100000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	var err error
	var n, m int

	scanner.Scan()
	n, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	data := stringToIntSlice(n, scanner.Text())

	scanner.Scan()
	m, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	searchFor := stringToIntSlice(m, scanner.Text())

	offset := 0
	ol := n - m + 1

Outer:
	for i := 0; i < ol; i++ {

		offset = searchFor[0] - data[i]
		for g := 1; g < m; g++ {
			if data[i+g]+offset != searchFor[g] {
				continue Outer
			}
		}
		s.WriteString(fmt.Sprint(i+1) + " ")
	}
}

func stringToIntSlice(l int, s string) []int {
	data := make([]int, 0, l)
	start := 0
	for i, r := range s {
		if !unicode.IsSpace(r) {
			continue
		}
		d, _ := strconv.Atoi(s[start:i])
		data = append(data, d)
		start = i + 1
	}
	d, _ := strconv.Atoi(s[start:])
	data = append(data, d)
	return data
}

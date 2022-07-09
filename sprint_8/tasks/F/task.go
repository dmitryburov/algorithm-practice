package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

	countMap := map[string]int{}
	for i := 0; i < n; i++ {
		scanner.Scan()
		str := scanner.Text()

		if unicode.IsSpace(rune(str[len(str)-1])) {
			str = str[:len(str)-1]
		}
		countMap[str]++
	}

	max := -1
	var maxWords []string = nil
	for k, v := range countMap {
		if max == -1 || v > max {
			max = v
			maxWords = []string{k}
			continue
		} else if v == max {
			maxWords = append(maxWords, k)
		}
	}

	sort.Strings(maxWords)
	s.WriteString(fmt.Sprint(maxWords[0]))
}

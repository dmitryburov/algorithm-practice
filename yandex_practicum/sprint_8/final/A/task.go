package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Stack [][]byte

const (
	BRACKET_OPEN  = byte('[')
	BRACKET_CLOSE = byte(']')
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

// Solution решение задачи
func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)

	const bufCapacity = 10000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	var prefix, current string

	for i := 0; i < n; i++ {
		scanner.Scan()
		if i == 0 {
			prefix = unpack(scanner.Bytes())
			continue
		}

		current = unpack(scanner.Bytes())
		prefix = current[:findIndex(prefix, current)]
	}

	s.WriteString(prefix)
}

// unpack распаковка строки
func unpack(s []byte) string {
	var stack Stack
	var builder strings.Builder
	var top, curr []byte

	for i := 0; i < len(s); i++ {
		// fix
		if unicode.IsSpace(rune(s[i])) {
			continue
		}

		switch s[i] {
		case BRACKET_CLOSE:
			top = stack.Pop()
			curr = nil

			for !unicode.IsDigit(rune(top[0])) {
				curr = append(top, curr...)
				top = stack.Pop()
			}

			if num, err := strconv.Atoi(string(top)); err != nil {
				log.Fatal(err)
			} else {
				stack.Push(bytes.Repeat(curr, num))
			}
		case BRACKET_OPEN:
			continue
		default:
			stack.Push([]byte(string(s[i])))
		}
	}

	for i := 0; i < len(stack); i++ {
		builder.Write(stack[i])
	}

	return builder.String()
}

// findIndex поиск первого неравного индекса
func findIndex(a, b string) int {
	var max int

	if len(a) > len(b) {
		max = len(a)
	} else {
		max = len(b)
	}

	var i int
	for ; i < max; i++ {
		if i >= len(a) || i >= len(b) || a[i] != b[i] {
			break
		}
	}

	return i
}

// IsEmpty пустой стек
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push добавление в стек
func (s *Stack) Push(str []byte) {
	*s = append(*s, str)
}

// Pop последний элемент из стека
func (s *Stack) Pop() []byte {
	if s.IsEmpty() {
		return nil
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

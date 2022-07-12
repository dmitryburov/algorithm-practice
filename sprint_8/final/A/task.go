package main

/*
--- Посылка
69359828

--- Принцип работы
Первая строка распаковывается как обычно, а последующие распаковываю до наибольшего возможного префикса.
Например, если при распаковке первой строки ее длина оказалась равна 6, а длинна второй оказалась, 500,
то нет никакого смысла распаковывать все 500 символов.

Распаковка идет в 2 этапа: сначала ищем начало запакованной строки с помощью проверки руны на число,
потом мы находим конец запакованной строки, и отправляем запакованную строку в следующий шаг рекурсии.

Дополнительные комментарии в коде.

--- Временная сложность
У нас есть все строки в количестве n, и в ней есть минимальная строка длинны m, которая является общим префиксом для всех строк,
то мы имеем на каждом шаге распаковку длинны m и сравнение со старым префиксом тоже длинны m.
Итого O(n * (m + m))

--- Пространственная сложность
Мы имеем текущий общий префикс длинны m и новую строку для проверки длинны m - O(m + m)

*/
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

type Stack []string

const (
	BRACKET_LEFT  = "["
	BRACKET_RIGHT = "]"
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

	lines := make([]string, n)

	// распаковываем строки
	for i := 0; i < n; i++ {
		scanner.Scan()
		lines[i] = unpack(scanner.Bytes())
	}

	s.WriteString(searchPrefix(n, lines))
}

// unpack распаковка строки
// убрал рекурсию и переделал на стек
func unpack(line []byte) string {
	var stack Stack
	var ch, top, curr string

	for i := 0; i < len(line); i++ {
		ch = string(line[i])

		switch ch {
		case BRACKET_RIGHT:
			top = stack.Pop()
			curr = ""

			for !unicode.IsDigit(rune(top[0])) {
				curr = top + curr
				top = stack.Pop()
			}

			if num, err := strconv.Atoi(top); err == nil {
				stack.Push(strings.Repeat(curr, num))
			}
		case BRACKET_LEFT:
			continue
		default:
			stack.Push(ch)
		}
	}

	result := ""
	for len(stack) > 0 {
		result = stack.Pop() + result
	}

	return result
}

// searchPrefix поиск префикса
func searchPrefix(n int, lines []string) string {
	tmp := lines[0]

	for i := 0; i < len(tmp); i++ {
		ch := tmp[i]
		for j := 1; j < n; j++ {
			line := lines[j]
			if i > len(line) || line[i] != ch {
				return tmp[:i]
			}
		}
	}

	return tmp
}

// IsEmpty пустой стек
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push добавление в стек
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop последний элемент из стека
func (s *Stack) Pop() string {
	if s.IsEmpty() {
		return ""
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element
	}
}

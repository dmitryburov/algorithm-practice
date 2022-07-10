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

	prefix := ""
	current := ""

	for i := 0; i < n; i++ {
		scanner.Scan()
		if i == 0 {
			prefix = unpack(scanner.Text())
			continue
		}
		current = lengthUnpack(scanner.Text(), len(prefix), 0)
		prefix = current[:firstIndex(prefix, current)]
	}

	s.WriteString(prefix)
}

// unpack распаковка строки
func unpack(s string) string {
	builder := strings.Builder{}

	for i := 0; i < len(s); {
		if unicode.IsDigit(rune(s[i])) {
			howMuch, _ := strconv.Atoi(string(s[i]))
			i += 2
			end := searchEnd(s, i)
			unpackedS := unpack(s[i:end])

			for g := 0; g < howMuch; g++ {
				builder.WriteString(unpackedS)
			}
			i = end + 1

		} else {
			builder.WriteByte(s[i])
			i++
		}
	}

	return builder.String()
}

// firstIndex поиск первого неравного индекса
func firstIndex(a, b string) int {
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

// lengthUnpack
func lengthUnpack(s string, needLen int, currLen int) string {
	builder := strings.Builder{}
	var length int

	for i := 0; i < len(s); {
		length = builder.Len() + currLen
		if length >= needLen {
			return cutString(builder.String(), length-needLen)
		}

		if unicode.IsDigit(rune(s[i])) {
			howMuch, _ := strconv.Atoi(string(s[i]))
			i += 2
			end := searchEnd(s, i)
			unpackedS := lengthUnpack(s[i:end], needLen, builder.Len()+currLen)

			for g := 0; g < howMuch; g++ {
				length = builder.Len() + currLen
				if length >= needLen {
					return cutString(builder.String(), length-needLen)
				}

				builder.WriteString(unpackedS)
			}

			i = end + 1
			continue
		}

		builder.WriteByte(s[i])
		i++
	}

	return builder.String()
}

// cutString обрезка строки
func cutString(s string, to int) string {
	return s[:len(s)-to]
}

// searchEnd конец запакованной строки
func searchEnd(s string, start int) int {
	bracketCount := 1
	i := start

	for ; i < len(s); i++ {
		if s[i] == '[' {
			bracketCount++
			continue
		} else if s[i] == ']' {
			bracketCount--
			if bracketCount == 0 {
				break
			}
		}
	}

	return i
}

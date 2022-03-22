package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
)

// L. Лишняя буква
// Васе очень нравятся задачи про строки, поэтому он придумал свою.
// Есть 2 строки s и t, состоящие только из строчных букв.
// Строка t получена перемешиванием букв строки s и добавлением 1 буквы в случайную позицию. Нужно найти добавленную букву.
//
// Формат ввода:
// На вход подаются строки s и t, разделённые переносом строки. Длины строк не превосходят 1000 символов.
// Строки не бывают пустыми.
//
// Формат вывода:
// Выведите лишнюю букву.
//
// Примеры:
// inp: 	>>> 	out:
//    abcd				e
//    abcde
//
// inp: 	>>> 	out:
//    go				g
//	  ogg
//
// inp: 	>>> 	out:
//    xtkpx				c
//	  xkctpx
func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)

	var find string
	var values [][]string

	for i := 1; i <= 2; i++ {
		scanner.Scan()

		parse := strings.Split(scanner.Text(), "")
		sort.StringSlice(parse).Sort()

		values = append(values, parse)
	}

	for i := 0; i < len(values[1]); i++ {
		if i+1 > len(values[0]) {
			find = values[1][i]
			break
		}
		if values[1][i] != values[0][i] {
			find = values[1][i]
			break
		}
	}

	_, _ = writer.WriteString(find)
	_, _ = writer.WriteString("\n")
	_ = writer.Flush()
}

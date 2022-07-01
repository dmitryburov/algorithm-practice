package main

/*
Задачка показалась как-то легкой.
Я прочитал про метрику "Расстояние Левенштейна" - https://ru.wikipedia.org/wiki/%D0%A0%D0%B0%D1%81%D1%81%D1%82%D0%BE%D1%8F%D0%BD%D0%B8%D0%B5_%D0%9B%D0%B5%D0%B2%D0%B5%D0%BD%D1%88%D1%82%D0%B5%D0%B9%D0%BD%D0%B0
По сути, ответ там и был в псевдокоде, как его решать.

--- Посылка
69267239

--- Принцип работы
В Го удобно получить строку в байтах и с байтами быстрее работать.
Отсекая базовые случаи вначале собираем матрицу, а затем

В алгоритме используется двумерная динамика.
Дополнительные комментарии в коде.

--- Временная сложность
Сложность двумерной динамики - O(n*m)

--- Пространственная сложность
Мы храним только матрицу и 2 строки.
O(n*m) - если не учитывать строки

*/
import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

// Solution решение задачи
func Solution(r io.Reader, w *strings.Builder) {
	scanner := bufio.NewScanner(r)

	// считаем строки в байты (!)
	scanner.Scan()
	s := scanner.Bytes()
	scanner.Scan()
	t := scanner.Bytes()

	lenS, lenT := len(s), len(t)

	// fix базовые случаи
	if lenS == 0 || lenT == 0 || bytes.Compare(s, t) == 0 {
		w.WriteString("0")
		return
	}

	matrix := make([]int, lenS+1)
	for i := 1; i < len(matrix); i++ {
		matrix[i] = i
	}

	for i := 1; i <= lenT; i++ {
		prev := i
		for j := 1; j <= lenS; j++ {
			if t[i-1] != s[j-1] {
				matrix[j-1] = getMinimum(matrix[j]+1, getMinimum(matrix[j-1]+1, prev+1))
			}
			matrix[j-1], prev = prev, matrix[j-1]
		}
		matrix[lenS] = prev
	}

	// результат
	w.WriteString(fmt.Sprint(matrix[lenS]))
}

// getMinimum сравнение минимального числа
func getMinimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

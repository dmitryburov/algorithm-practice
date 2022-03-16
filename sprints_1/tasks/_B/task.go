package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// B. Чётные и нечётные числа
// Алла придумала такую онлайн-игру: игрок нажимает на кнопку, и на экране появляются три случайных числа.
// Если все три числа оказываются одной чётности, игрок выигрывает.
// Напишите программу, которая по трём числам определяет, выиграл игрок или нет.
//
// Формат ввода:
// В первой строке записаны три случайных целых числа a, b и c. Числа не превосходят 10(9) по модулю.
//
// Формат вывода:
// Выведите «WIN», если игрок выиграл, и «FAIL» в противном случае.
//
// Примеры:
// inp: 1 2 -3	>>	out: FAIL
// inp: 7 11 7	>>	out: WIN
// inp: 6 -2 0	>>	out: WIN
func main() {

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	line := scanner.Text()

	var result = "WIN"
	var values = strings.Split(line, " ")
	var last int

	for i := 0; i < len(values); i++ {
		num, _ := strconv.Atoi(values[i])
		if num%2 == 0 {
			if last == 1 {
				result = "FAIL"
				break
			}
			last = 2
		} else {
			if last == 2 {
				result = "FAIL"
				break
			}
			last = 1
		}
	}

	_, _ = writer.WriteString(result)
	_, _ = writer.WriteString("\n")
	_ = writer.Flush()
}

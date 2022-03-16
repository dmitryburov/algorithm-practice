package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// A. Значения функции
// Вася делает тест по математике: вычисляет значение функций в различных точках.
// Стоит отличная погода, и друзья зовут Васю гулять. Но мальчик решил сначала закончить тест и только после этого идти к друзьям.
//
// К сожалению, Вася пока не умеет программировать. Зато вы умеете. Помогите Васе написать код функции, вычисляющей y = ax2 + bx + c.
// Напишите программу, которая будет по коэффициентам a, b, c и числу x выводить значение функции в точке x.
//
// Формат ввода:
// На вход через пробел подаются числа a, x, b, c. В конце ввода находится перенос строки.
//
// Формат вывода:
// Выведите одно число — значение функции в точке x.
//
// Примеры:
// inp: -8 -5 -2 7	>>	out: -183
// inp: 8 2 9 -10	>>	out: 40
func main() {
	var a, b, c, x int

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	writer := bufio.NewWriter(os.Stdout)

	scanner.Scan()
	line := scanner.Text()

	values := strings.Split(line, " ")

	a, _ = strconv.Atoi(values[0])
	x, _ = strconv.Atoi(values[1])
	b, _ = strconv.Atoi(values[2])
	c, _ = strconv.Atoi(values[3])

	result := a*(x*x) + b*x + c

	_, _ = writer.WriteString(strconv.Itoa(result))
	_, _ = writer.WriteString("\n")
	_ = writer.Flush()
}

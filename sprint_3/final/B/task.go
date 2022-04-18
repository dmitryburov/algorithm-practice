package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
ID посылки: 67500221

ПРИНЦИП РАБОТЫ ---
Мы делим масив на меньшие части и потом эти части рекурсивно переупорядовачиваем.
Комментарии в коде.

ВРЕМЕННАЯ СЛОЖНОСТЬ --
Алгоритм работает за O(n*log n), тк изначально мы проходим по всему массиву,
а затем мы повторяем процесс для меньших частей массива.

ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Исходный массив состоит из n элементов и занимает О(n) помаяти.
Поскольку условие задачи ограничивает помять в O(n), то памяти больше не выделяется.
Дополнительные затраты памяти выходят только на рекурсию и "компаратор".

*/

// User пользователь с параметрами
type User struct {
	Login   string
	Success int
	Fail    int
}

func main() {
	n, users, err := getInputData()
	if err != nil {
		showError(err)
	}

	// если больше 1 для расчета :)
	if len(users) > 1 {
		// сортируем
		quickSort(users, 0, len(users)-1)
	}

	// выводим
	for i := 0; i < n; i++ {
		fmt.Println(users[i].Login)
	}
}

// quickSort сортиравка массива
func quickSort(arr []User, left, right int) {

	// базовый случай
	if left >= right {
		return
	}

	pivot := sortPart(arr, left, right, comparator)

	quickSort(arr, left, pivot-1)
	quickSort(arr, pivot+1, right)
}

// sortPart точка опоры и сортировка части массива
func sortPart(arr []User, left, right int, compare func(a, b User) bool) int {
	p := arr[left]
	l := left + 1
	h := right

	for {
		for l <= h && compare(arr[h], p) {
			h--
		}

		for l <= h && !compare(arr[l], p) {
			l++
		}

		if l <= h {
			arr[l], arr[h] = arr[h], arr[l]
		} else {
			break
		}
	}

	arr[left], arr[h] = arr[h], arr[left]
	return h
}

// comparator сравнение значений (компаратор)
func comparator(a, b User) bool {

	// если равенство числа решённых задач
	if a.Success == b.Success {

		// если и в штрафах равенство - выводим по лексике
		if a.Fail == b.Fail {
			if strings.Compare(a.Login, b.Login) > 0 {
				return true
			}
			return false
		}

		// сравним по штрафам
		return a.Fail > b.Fail
	}

	// сравним по решенным
	return a.Success < b.Success
}

// getInputData подготовка входных данных
func getInputData() (n int, users []User, err error) {
	var input *os.File
	var bufStr string

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	bufStr = scanner.Text()
	n, err = strconv.Atoi(bufStr)
	if err != nil {
		return
	}

	users = make([]User, n)
	var user User
	for i := 0; i < n; i++ {
		scanner.Scan()
		bufStr = scanner.Text()

		strArr := strings.Split(bufStr, " ")
		user = User{
			Login: strArr[0],
		}

		if strArr[1] != "0" {
			user.Success, err = strconv.Atoi(strArr[1])
		}
		if strArr[2] != "0" {
			user.Fail, err = strconv.Atoi(strArr[2])
		}

		users[i] = user
	}

	return
}

// getInputFromFile получение ввода из файла
func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

// showError вывод ошибки
func showError(err interface{}) {
	panic(err)
}

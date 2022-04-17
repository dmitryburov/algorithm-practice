package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
ID посылки:

ПРИНЦИП РАБОТЫ ---

ВРЕМЕННАЯ СЛОЖНОСТЬ --

ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --

*/

// User пользователь с параметрами
type User struct {
	Login         string
	Success, Fail int
}

func main() {
	n, users, err := getInputData()
	if err != nil {
		showError(err)
	}

	// сортируем
	quickSort(users, 0, len(users)-1)

	// выводим
	for i := 0; i < n; i++ {
		fmt.Println(users[i].Login)
	}
}

// quickSort сортиравка массива
func quickSort(arr []User, left, right int) {

}

// compare сравнение значений (компаратор)
func compare(a, b int) bool {
	return a > b
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

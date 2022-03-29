package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Stack struct {
	items []float64
}

func main() {

	// получаем данные
	operation, err := getInputData()
	if err != nil {
		showError(err)
	}

	// решение
	result, err := solution(operation)
	if err != nil {
		showError(err)
	}

	fmt.Println(result)
}

// solution решение задачи
func solution(operation string) (float64, error) {
	var (
		err   error
		stack = &Stack{}
		items = strings.Split(operation, " ")
		num   float64
	)

	for i := range items {
		if items[i] == "+" || // не совсем лаконичное решение :)
			items[i] == "-" ||
			items[i] == "*" ||
			items[i] == "/" {
			if err = stack.calculate(items[i]); err != nil {
				return 0, err
			}
		} else {
			num, err = strconv.ParseFloat(items[i], 64)
			if err != nil {
				return 0, err
			}
			stack.push(num)
		}
	}

	return stack.peak()
}

// calculate производит операцию
func (s *Stack) calculate(operation string) error {

	var firstNum, secondNum float64
	var value float64
	var err error

	secondNum, err = s.pop()
	if err != nil {
		return err
	}

	firstNum, err = s.pop()
	if err != nil {
		return err
	}

	switch operation {
	case "+":
		value = firstNum + secondNum
		break
	case "-":
		value = firstNum - secondNum
		break
	case "*":
		value = firstNum * secondNum
		break
	case "/":
		value = math.Floor(firstNum / secondNum)
		break
	default:
		break
	}

	s.push(value)
	return nil
}

// pop добавляет элемент
func (s *Stack) pop() (float64, error) {
	if len(s.items) > 0 {
		x := s.items[len(s.items)-1]
		s.items = s.items[:len(s.items)-1]
		return x, nil
	}

	return 0, fmt.Errorf("stack is empty")
}

// push добавляет элемент
func (s *Stack) push(value float64) {
	s.items = append(s.items, value)
}

// peak получает последний элемент
func (s *Stack) peak() (float64, error) {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1], nil
	}

	return 0, fmt.Errorf("empty last value")
}

// getInputData парсинг входных данных
func getInputData() (operation string, err error) {

	input, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	reader := bufio.NewReader(input)

	strNum, _, _ := reader.ReadLine()
	operation = string(strNum)

	// clear bufio
	defer reader.Reset(reader)
	return
}

// getInputFromFile получение input из файла
func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

// showError обработка ошибки
func showError(err interface{}) {
	panic(err)
}

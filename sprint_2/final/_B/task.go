package main

// ID посылки: 66613162

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	ErrInputValue = "not found last value"

	ActionPlus     = "+"
	ActionMinus    = "-"
	ActionDivision = "/"
	ActionMulti    = "*"
)

// Node элемент списка
type Node struct {
	value float64
	next  *Node
}

// NodeList список
type NodeList struct {
	head *Node
	size int
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
		list  = &NodeList{}
		items = strings.Split(operation, " ")
		num   float64
	)

	for i := range items {
		if items[i] == ActionPlus ||
			items[i] == ActionMinus ||
			items[i] == ActionDivision ||
			items[i] == ActionMulti {
			if err = list.calculate(items[i]); err != nil {
				return 0, err
			}
		} else {
			num, err = strconv.ParseFloat(items[i], 64)
			if err != nil {
				return 0, err
			}
			list.push(num)
		}
	}

	return list.peak()
}

// calculate производит операцию
func (n *NodeList) calculate(operation string) error {

	var firstNum, secondNum float64
	var value float64
	var err error

	secondNum, err = n.pop()
	if err != nil {
		return err
	}

	firstNum, err = n.pop()
	if err != nil {
		return err
	}

	switch operation {
	case ActionPlus:
		value = firstNum + secondNum
		break
	case ActionMinus:
		value = firstNum - secondNum
		break
	case ActionMulti:
		value = firstNum * secondNum
		break
	case ActionDivision:
		value = math.Floor(firstNum / secondNum)
		break
	default:
		break
	}

	n.push(value)
	return nil
}

// push добавляет элемент в список
func (n *NodeList) push(value float64) {
	node := Node{
		value: value,
		next:  nil,
	}

	if n.head == nil {
		n.head = &node
	} else {
		last := n.head
		n.head = &node
		n.head.next = last
	}
	n.size++
}

// pop извлекает элемент из списка
func (n *NodeList) pop() (float64, error) {
	if !n.isEmpty() {
		x := n.head.value
		head := n.head.next
		n.head = head

		n.size--
		return x, nil
	}

	return 0, fmt.Errorf(ErrInputValue)
}

// peak получает последний элемент
func (n *NodeList) peak() (float64, error) {
	if !n.isEmpty() {
		return n.head.value, nil
	}

	return 0, fmt.Errorf(ErrInputValue)
}

// isEmpty проверяет пуст ли список
func (n *NodeList) isEmpty() bool {
	return n.size == 0
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

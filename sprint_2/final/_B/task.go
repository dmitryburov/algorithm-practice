package main

// ID посылки: 66643660

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

// Action таблица операций
type Action map[string]func(a, b float64) float64

var actions Action

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
		items = strings.Split(operation, " ")
		num   float64
	)

	list := &NodeList{}
	actions.initActions()

	for i := range items {
		if actions.isAction(items[i]) {
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
	var err error

	secondNum, err = n.pop()
	if err != nil {
		return err
	}

	firstNum, err = n.pop()
	if err != nil {
		return err
	}

	n.push(actions[operation](firstNum, secondNum))
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

// initActions
func (Action) initActions() {
	actions = make(map[string]func(a, b float64) float64)
	actions[ActionPlus] = func(a, b float64) float64 {
		return a + b
	}
	actions[ActionMinus] = func(a, b float64) float64 {
		return a - b
	}
	actions[ActionMulti] = func(a, b float64) float64 {
		return a * b
	}
	actions[ActionDivision] = func(a, b float64) float64 {
		return math.Floor(a / b)
	}
}

// isAction проряет является ил строка операцией
func (Action) isAction(item string) bool {
	if _, ok := actions[item]; ok {
		return true
	}

	return false
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

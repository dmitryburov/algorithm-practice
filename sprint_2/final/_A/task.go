package main

// ID посылки: 66641053

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Deque двусторонняя очередь
type Deque struct {
	stack []int
	head,
	tail,
	maxSize,
	size int
}

// command команда действия
type command struct {
	action string
	num    int
}

const (
	ErrStackEmpty = "error"
	ErrStackMax   = "error"

	ActionPushBack  = "push_back"
	ActionPushFront = "push_front"
	ActionPopFront  = "pop_front"
	ActionPopBack   = "pop_back"
)

func main() {

	// сбор данных из input
	q, commands, err := getInputData()
	if err != nil {
		showError(err)
	}

	// решение
	result := q.Run(commands)
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

// Run выполнение алгоритма
func (q *Deque) Run(commands []command) []string {
	var (
		err    error
		result []string
	)

	for i := 0; i < len(commands); i++ {
		switch commands[i].action {
		case ActionPushBack, ActionPushFront:
			if err = q.push(commands[i].num, commands[i].action); err != nil {
				result = append(result, err.Error())
			}
			break
		case ActionPopFront, ActionPopBack:
			result = append(result, q.pop(commands[i].action))
		default:
			break
		}
	}

	return result
}

// push добавление элемента
func (q *Deque) push(value int, action string) error {
	if q.isMax() {
		return fmt.Errorf(ErrStackMax)
	}

	if action == ActionPushBack {
		q.setStack(q.tail, value)
		q.tail = (q.tail + 1) % q.maxSize
	} else {
		q.setStack(q.head-1, value)
		q.head = (q.head - 1) % q.maxSize
	}

	q.size++
	return nil
}

// pop вывести элемент
func (q *Deque) pop(action string) string {
	if q.isEmpty() {
		return ErrStackEmpty
	}

	var x int

	if action == ActionPopBack {
		x = q.stack[q.getIndex(q.tail-1)]
		q.setStack(q.tail-1, 0)
		q.tail = (q.tail - 1) % q.maxSize
	} else {
		x = q.stack[q.getIndex(q.head)]
		q.setStack(q.head, 0)
		q.head = (q.head + 1) % q.maxSize
	}

	q.size--
	return fmt.Sprint(x)
}

// setStack меняет значение в стеке
func (q *Deque) setStack(index, value int) {
	q.stack[q.getIndex(index)] = value
}

// getIndex перевод "обратных" индексов
func (q *Deque) getIndex(index int) int {
	if index < 0 {
		return len(q.stack) + index
	}

	return index
}

// isEmpty определяет, пуст ли дек
func (q *Deque) isEmpty() bool {
	return q.size == 0
}

// isMax определяет, переполнен ли дек
func (q *Deque) isMax() bool {
	return q.size >= q.maxSize
}

// getInputData парсинг входных данных
func getInputData() (q *Deque, commands []command, err error) {

	input, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	reader := bufio.NewReader(input)

	var n, m int
	strNum, _, _ := reader.ReadLine()
	n, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	strNum, _, _ = reader.ReadLine()
	m, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	q = &Deque{
		stack:   make([]int, m),
		maxSize: m,
	}

	var com command
	commands = make([]command, 0, n)

	for i := 0; i < n; i++ {
		strNums, _ := reader.ReadString('\n')
		comStr := strings.Split(strings.TrimSpace(strNums), " ")

		com = command{
			action: comStr[0],
		}
		if len(comStr) == 2 {
			com.num, _ = strconv.Atoi(comStr[1])
		}

		commands = append(commands, com)
	}

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

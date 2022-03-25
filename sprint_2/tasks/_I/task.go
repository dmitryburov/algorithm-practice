package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	Name string
	Step int
}

type queue struct {
	stack    []int
	commands []command
	max      int
}

func main() {
	n, q, err := getInputData()
	if err != nil {
		showError(err)
	}

	// решение
	result := solution(n, q)
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
	}
}

func solution(n int, q *queue) (result []string) {
	for i := 0; i < n; i++ {
		switch q.commands[i].Name {
		case "size":
			result = append(result, q.size())
			break
		case "push":
			if len(q.stack) >= q.max {
				result = append(result, "error")
			} else {
				q.push(q.commands[i].Step)
			}
			break
		case "pop":
			result = append(result, q.pop())
			break
		case "peek":
			result = append(result, q.peek())
			break
		default:
			break
		}
	}

	return result
}

// push добавить число x в очередь
func (q *queue) push(x int) {
	q.stack = append(q.stack, x)
}

// pop удалить число из очереди и вывести на печать
func (q *queue) pop() (result string) {
	if q.isEmpty() {
		result = "None"
	} else {
		result = fmt.Sprint(q.stack[0])
		q.stack = q.stack[1:]
	}

	return
}

// peek напечатать первое число в очереди
func (q *queue) peek() (result string) {
	if q.isEmpty() {
		result = "None"
	} else {
		result = fmt.Sprint(q.stack[0])
	}

	return
}

// isEmpty проверка пустой очереди
func (q *queue) isEmpty() bool {
	return len(q.stack) == 0
}

// size вернуть размер очереди
func (q *queue) size() string {
	return fmt.Sprint(len(q.stack))
}

// getInputData парсинг входных данных
func getInputData() (n int, q *queue, err error) {

	inp, err := getInputFromFile()
	if err != nil {
		return
	}
	defer func(inp *os.File) {
		_ = inp.Close()
	}(inp)

	reader := bufio.NewReader(inp)

	strNum, _, _ := reader.ReadLine()
	n, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	var k int
	strNum, _, _ = reader.ReadLine()
	k, err = strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	q = &queue{
		stack: make([]int, 0, k),
		max:   k,
	}

	q.commands = make([]command, 0, n)

	var num int
	var com command
	for i := 0; i < n; i++ {
		strNums, _ := reader.ReadString('\n')
		comStr := strings.Split(strings.TrimSpace(strNums), " ")

		com = command{
			Name: comStr[0],
		}
		if len(comStr) == 2 {
			num, _ = strconv.Atoi(comStr[1])
			com.Step = num
		}

		q.commands = append(q.commands, com)
	}

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

// showError вывод ошибки
func showError(err error) {
	panic(err)
}

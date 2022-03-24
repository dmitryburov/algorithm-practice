package main

// Зачада G. Стек - MaxEffective
// ID посылки:

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type queue struct {
	stack,
	max []int
}
type command struct {
	Name string
	Step int
}

func main() {
	n, commands, err := getInputData()
	if err != nil {
		showError(err)
	}

	// решение
	result := solution(n, commands)
	fmt.Println(strings.Join(result, "\n"))
}

// solution решение задачи
func solution(n int, commands []command) []string {
	var result []string
	var q = &queue{
		stack: []int{},
		max:   []int{},
	}

	for i := 0; i < n; i++ {
		switch commands[i].Name {
		case "get_max":
			if q.isEmpty() {
				result = append(result, "None")
			} else {
				result = append(result, q.getMax())
			}
			break
		case "push":
			q.push(commands[i].Step)
			break
		case "pop":
			if q.isEmpty() {
				result = append(result, "error")
			} else {
				q.pop()
			}
			break
		default:
			break
		}
	}

	return result
}

// push добавить число в стек
func (q *queue) push(num int) {
	q.stack = append(q.stack, num)
}

// pop удалить число с вершины стека
func (q *queue) pop() {
	q.stack = q.stack[:len(q.stack)-1]
}

// max максимальное число в стеке
func (q *queue) getMax() string {
	return fmt.Sprint(q.max[len(q.max)-1])
}

// empty проверка на пустоту
func (q *queue) isEmpty() bool {
	return len(q.stack) == 0
}

// getInputData парсинг входных данных
func getInputData() (n int, commands []command, err error) {

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

	commands = make([]command, 0, n+1)

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

		commands = append(commands, com)
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

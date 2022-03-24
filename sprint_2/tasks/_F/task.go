package main

// Зачада F. Стек - Max
// ID посылки: 66373456

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

func solution(n int, commands []command) []string {
	var stack []int
	var result []string

	getMaxStack := func() string {
		if len(stack) == 0 {
			return "None"
		}

		var m int
		for i, v := range stack {
			if i == 0 || v > m {
				m = v
			}
		}
		return fmt.Sprint(m)
	}

	for i := 0; i < n; i++ {

		switch commands[i].Name {
		case "get_max":
			result = append(result, getMaxStack())
			break
		case "push":
			stack = append(stack, commands[i].Step)
			break
		case "pop":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else {
				result = append(result, "error")
			}
			break
		default:
			break
		}
	}

	return result
}

func main() {
	n, commands, err := getInputData()
	if err != nil {
		showError(err)
	}

	// решение
	result := solution(n, commands)
	if len(result) > 0 {
		for i := 0; i < len(result); i++ {
			fmt.Println(result[i])
		}
	}
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

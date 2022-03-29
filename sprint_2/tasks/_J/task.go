package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	action string
	num    int
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

func solution(n int, commands []command) (result []string) {
	var stack []int

	for i := 0; i < n; i++ {
		switch commands[i].action {
		case "get":
			if len(stack) == 0 {
				result = append(result, "error")
			} else {
				result = append(result, fmt.Sprint(stack[0]))
				stack = stack[1:]
			}
			break
		case "put":
			stack = append(stack, commands[i].num)
			break
		case "size":
			result = append(result, fmt.Sprint(len(stack)))
			break
		default:
			break
		}
	}

	return
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

	commands = make([]command, 0, n)

	var num int
	var com command
	for i := 0; i < n; i++ {
		strNums, _ := reader.ReadString('\n')
		comStr := strings.Split(strings.TrimSpace(strNums), " ")

		com = command{
			action: comStr[0],
		}
		if len(comStr) == 2 {
			num, _ = strconv.Atoi(comStr[1])
			com.num = num
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
func showError(err interface{}) {
	panic(err)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	if idx == 0 {
		head = head.next
	}

	if head != nil {
		fmt.Println(head.data)
	}

	if head != nil && head.next != nil {
		return Solution(head.next, idx-1)
	}
	return nil
}

func main() {
	inp, err := getInputFromFile()
	if err != nil {
		showError(err)
	}
	defer inp.Close()

	reader := bufio.NewReader(inp)

	strNum, _, _ := reader.ReadLine()
	n, err := strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	var lastNode *ListNode
	for i := 0; i < n; i++ {
		list, _ := reader.ReadString('\n')
		node := ListNode{
			data: strings.TrimSpace(list),
			next: lastNode,
		}
		lastNode = &node
	}

	strNum, _, _ = reader.ReadLine()
	i, err := strconv.Atoi(string(strNum))
	if err != nil {
		return
	}

	Solution(lastNode, i)
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

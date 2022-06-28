package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
В решении я использовал алгоритм Дийкстры для построения максимального остового дерева и
приоритетную очередь для хранения ребер. Можно сказать почти "скопировано" из теории ))

--- Посылка


--- Принцип работы


--- Временная сложность


--- Пространственная сложность


*/

const ErrOops = "Oops! I did it again"

func main() {
	res := strings.Builder{}
	Solution(os.Stdin, &res)

	fmt.Println(res.String())
}

// Solution решение задачи
func Solution(r io.Reader, w *strings.Builder) {
	var err error
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	// количество вершин и ребер
	var n, m int
	nmData := strings.Fields(scanner.Text())
	n, err = strconv.Atoi(nmData[0])
	if err != nil {
		log.Fatal(err)
	}
	m, err = strconv.Atoi(nmData[1])
	if err != nil {
		log.Fatal(err)
	}

}

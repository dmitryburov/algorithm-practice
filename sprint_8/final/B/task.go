package main

/*
--- Посылка
69358298

--- Принцип работы
Алгоритм решается с помощью структура Бора и динамического программирования.

При обходе строки мы складываем в массив построки, которые привели к текущему символу (ведут ли какие-то ноды из dp[i - 1)).
Если в массиве есть ноды, которые являются концами подстрок, то мы ищем символ в рутовой ноде.
Если какой-то из символов подстрок, который пришел к концу строки является концом подстроки, то тогда ответ "YES", если нет, то "NO"

Дополнительные комментарии в коде.

--- Временная сложность
В теории сложность формирования Бора это O(l), где l — суммарная длина слов во множестве.
Высчитывая результат на основе ДП это по сложности будет O(s), где s - длина строки.
Поскольку итерация проходит n нод по каждому символу, то сложность будет O(l + (s^2))

--- Пространственная сложность
Мы храним структуру Бора и массив подстрок - O(l + s). Но что-то не уверен в правильности))

*/
import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	SUCCESS = "YES"
	FAIL    = "NO"
)

type Node struct {
	isBlack bool
	tree    PrefixMap
}

type PrefixMap map[string]*Node

type DP [][]*Node

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

// Solution решение задачи
func Solution(r io.Reader, s *strings.Builder) {

	t, rootNode := initInput(bufio.NewScanner(r))

	if rootNode.tree[string(t[0])] == nil {
		s.WriteString(FAIL)
		return
	}

	dp := make(DP, len(t))
	dp[0] = append(dp[0], rootNode.tree[string(t[0])])

	for i := 1; i < len(t); i++ {
		c := string(t[i])
		var isSomeBlack bool
		dp[i], isSomeBlack = findNodes(dp[i-1], c)

		if isSomeBlack && rootNode.tree[c] != nil {
			dp[i] = append(dp[i], rootNode.tree[c])
		}

		if len(dp[i]) == 0 {
			break
		}
	}

	for _, v := range dp[len(t)-1] {
		if v.isBlack {
			s.WriteString(SUCCESS)
			return
		}
	}

	s.WriteString(FAIL)
}

// findNodes поиск ноды по символу
func findNodes(n []*Node, char string) ([]*Node, bool) {
	var res []*Node
	var isBlack bool

	for _, v := range n {
		if v.isBlack {
			isBlack = true
		}
		if v.tree[char] != nil {
			res = append(res, v.tree[char])
		}
	}

	if len(res) == 0 {
		return res, isBlack
	}

	return res, isBlack
}

// addPrefix добавление префикса в ноду
func addPrefix(root *Node, s string) {
	for i := 0; i < len(s); i++ {
		if _, ok := root.tree[string(s[i])]; !ok {
			root.tree[string(s[i])] = &Node{
				isBlack: false,
				tree:    PrefixMap{},
			}
		}

		root = root.tree[string(s[i])]
	}

	root.isBlack = true
}

// initInput парсим input-данные
func initInput(scanner *bufio.Scanner) (t string, root *Node) {
	const bufCapacity = 10000000 // fix long string
	buf := make([]byte, bufCapacity)
	scanner.Buffer(buf, bufCapacity)

	scanner.Scan()
	t = scanner.Text()

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	root = &Node{
		isBlack: false,
		tree:    PrefixMap{},
	}

	for i := 0; i < n; i++ {
		scanner.Scan()
		p := scanner.Text()
		if unicode.IsSpace(rune(p[len(p)-1])) {
			addPrefix(root, p[:len(p)-1])
		} else {
			addPrefix(root, p)
		}
	}

	return
}

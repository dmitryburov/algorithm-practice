package main

/*
ID посылки:

Работа алгоритма:

Временная сложность:
В лучшем случае все операции происходят за О(1), в худшем случае за O(n)

Пространственная сложность:


*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	ActionGet    = "get"
	ActionPut    = "put"
	ActionDelete = "delete"
	ActionNone   = "None"
)

var (
	minLoadFactor    = 0.25
	maxLoadFactor    = 0.75
	defaultTableSize = 3
)

// Node односвязный список
type Node struct {
	key,
	value string
	next *Node
}

// HashTable хеш таблица
type HashTable struct {
	buckets []*Node
	size    int
}

// HashCommand список команд из input
type HashCommand struct {
	Action,
	Key,
	Value string
}

// NewHashTable инициализация новой таблицы
// таблица с определенным размером (тк не требуется рехеширование и расширение таблицы)
func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*Node, size),
		size:    size,
	}
}

func main() {
	data, err := getInputData()
	if err != nil {
		showError(err)
	}

	h := NewHashTable(len(data))

	for i := 0; i < len(data)-1; i++ {
		switch data[i].Action {
		case ActionGet:
			fmt.Println(ActionGet, data[i].Key)
			if value, ok := h.GetKey(data[i].Key); ok {
				fmt.Println(value)
			} else {
				fmt.Println(ActionNone)
			}
			break
		case ActionPut:
			fmt.Println(ActionPut, data[i].Key, data[i].Value)
			h.PutKey(data[i].Key, data[i].Value)
			break
		case ActionDelete:
			fmt.Println(ActionDelete, data[i].Key)
			if value, ok := h.DeleteKey(data[i].Key); ok {
				fmt.Println(value)
			} else {
				fmt.Println(ActionNone)
			}
			break
		default:
			break
		}
	}
}

// PutKey добавление пары ключ-значение.
// Если заданный ключ уже есть в таблице, то соответствующее ему значение обновляется.
func (h *HashTable) PutKey(key, value string) {
	index := h.getIndex(key)

	if h.buckets[index] == nil {
		h.buckets[index] = &Node{key: key, value: value}
	} else {
		startingNode := h.buckets[index]
		for ; startingNode.next != nil; startingNode = startingNode.next {
			if startingNode.key == key {
				startingNode.value = value
				return
			}
		}
		startingNode.next = &Node{key: key, value: value}
	}
}

// GetKey получение значения по ключу.
// Если ключа нет в таблице, то вывести «None». Иначе вывести найденное значение.
func (h *HashTable) GetKey(key string) (string, bool) {
	index := h.getIndex(key)
	if h.buckets[index] != nil {
		startingNode := h.buckets[index]
		for ; ; startingNode = startingNode.next {
			if startingNode.key == key {
				return startingNode.value, true
			}
			if startingNode.next == nil {
				break
			}
		}
	}
	return "", false
}

// DeleteKey удаление ключа из таблицы.
// Если такого ключа нет, то вывести «None», иначе вывести хранимое по данному ключу значение и удалить ключ.
func (h *HashTable) DeleteKey(key string) (string, bool) {
	return "", false
}

// hash функция хеширования Дженкинса
// источник https://en.wikipedia.org/wiki/Jenkins_hash_function
func (h *HashTable) hash(key string) (hash uint32) {
	hash = 0
	for _, ch := range key {
		hash += uint32(ch)
		hash += hash << 10
		hash ^= hash >> 6
	}

	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15

	return
}

// getIndex
func (h *HashTable) getIndex(key string) (index int) {
	return int(h.hash(key)) % h.size
}

// getInputData подготовка входных данных
func getInputData() (data []HashCommand, err error) {
	var input *os.File
	var bufStr string
	var n int

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	bufStr = scanner.Text()
	n, err = strconv.Atoi(bufStr)
	if err != nil {
		return
	}

	data = make([]HashCommand, n)
	var arrStr []string
	for i := 0; i < n; i++ {
		scanner.Scan()
		arrStr = strings.Split(scanner.Text(), " ")
		data[i].Action = arrStr[0]
		data[i].Key = arrStr[1]

		if len(arrStr) == 3 {
			data[i].Value = arrStr[2]
		}
	}

	return
}

// getInputFromFile получение ввода из файла
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

package main

/*
ID посылки: 68020093

Принцип работы:
Хеш-таблица сопоставляют уникальный ключ со значением, чтобы обеспечить производительность поиска O(1) в среднем.
Коллизии решаются с помощью связанного списка (методом цепочек).
Функция хеширования использует метод Дженкинса и остаток деления по модулю.
Код подробно прокомментировал ниже.

Временная сложность:
В лучшем случае средняя сложность операций (put, get, delete) в хеш-таблице равняется О(1),
в худшем - за O(n), тк если по хешу элементы попали в односвязный список - нам придется прогуляться по нему.
Остальное время занимает вычисление хеша.

Пространственная сложность:
При инициализации хеш-таблицы мы заранее выделяем память на O(n) элементов.
При решении коллизий используются списки, они хранятся каждый в своей зоне памяти и требуют O(k) памяти,
где k - количетво элементов списка индекса хеша.

P.S. никак не получается корректно объяснить сложности алгоритмов, я старался)))
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

// Node односвязный список
type Node struct {
	key,
	value string
	next *Node
}

// HashTable хеш таблица
type HashTable struct {
	buckets []*Node
	size,
	n int
}

func main() {
	var n int
	var str string

	// получаем input файлом
	data, err := getInputData()
	if err != nil {
		showError(err)
	}

	// считываем первое n-число
	data.Scan()
	str = data.Text()
	n, err = strconv.Atoi(str)
	if err != nil {
		showError(err)
	}

	// инициализируем хэш-таблицу
	h := NewHashTable(n)

	// перебираем следующие n-команд
	var arrStr []string
	for i := 0; i < n; i++ {
		data.Scan()
		arrStr = strings.Split(data.Text(), " ")

		switch arrStr[0] {
		case ActionGet:
			if value, ok := h.GetKey(arrStr[1]); ok {
				fmt.Println(value)
			} else {
				fmt.Println(ActionNone)
			}
			break
		case ActionPut:
			h.PutKey(arrStr[1], arrStr[2])
			break
		case ActionDelete:
			if value, ok := h.DeleteKey(arrStr[1]); ok {
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

// testMain функция для тестирования
func testMain(n int, input []string) (result []string) {
	h := NewHashTable(n)
	var arrStr []string
	for i := 0; i < n; i++ {
		arrStr = strings.Split(input[i], " ")

		switch arrStr[0] {
		case ActionGet:
			if value, ok := h.GetKey(arrStr[1]); ok {
				result = append(result, value)
			} else {
				result = append(result, ActionNone)
			}
			break
		case ActionPut:
			h.PutKey(arrStr[1], arrStr[2])
			break
		case ActionDelete:
			if value, ok := h.DeleteKey(arrStr[1]); ok {
				result = append(result, value)
			} else {
				result = append(result, ActionNone)
			}
			break
		default:
			break
		}
	}

	return
}

// NewHashTable инициализация новой таблицы
// таблица с определенным размером (тк рехеширование и масштабирование хеш-таблицы не требуется)
func NewHashTable(size int) *HashTable {
	return &HashTable{
		buckets: make([]*Node, size),
		n:       size,
	}
}

// PutKey добавление пары ключ-значение.
// Если заданный ключ уже есть в таблице, то соответствующее ему значение обновляется.
func (h *HashTable) PutKey(key, value string) bool {
	index := h.getIndex(key)
	iterator := h.buckets[index]
	node := &Node{key, value, nil}

	if iterator == nil {
		h.buckets[index] = node
	} else {
		prev := &Node{}
		for iterator != nil {
			if iterator.key == key {
				iterator.value = value
				return false
			}
			prev = iterator
			iterator = iterator.next
		}
		prev.next = node
	}

	h.size++
	return true
}

// GetKey получение значения по ключу.
// Если ключа нет в таблице, то вывести «None». Иначе вывести найденное значение.
func (h *HashTable) GetKey(key string) (string, bool) {
	index := h.getIndex(key)
	iterator := h.buckets[index]

	for iterator != nil {
		if iterator.key == key {
			return iterator.value, true
		}
		iterator = iterator.next
	}

	return "", false
}

// DeleteKey удаление ключа из таблицы.
// Если такого ключа нет, то вывести «None», иначе вывести хранимое по данному ключу значение и удалить ключ.
func (h *HashTable) DeleteKey(key string) (val string, find bool) {
	index := h.getIndex(key)
	iterator := h.buckets[index]

	if iterator == nil {
		return
	}
	if iterator.key == key {
		val = iterator.value
		find = true
		h.buckets[index] = iterator.next
		h.size--
	} else {
		prev := iterator
		iterator = iterator.next
		for iterator != nil {
			if iterator.key == key {
				val = iterator.value
				find = true
				prev.next = iterator.next
				h.size--
				return
			}
			prev = iterator
			iterator = iterator.next
		}
	}

	return
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

// getIndex получение хешированного индекса
func (h *HashTable) getIndex(key string) int {
	return int(h.hash(key)) % h.n
}

// getInputData подготовка входных данных
func getInputData() (scan *bufio.Scanner, err error) {
	var input *os.File

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	return scanner, nil
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

package main

/*

ID посылки: 68060956

Принцип работы:
Алгоритм состоит в основном из 2х этапов:
1й - сбор индекса слов по документам, для этого создается мапа с индексом слов с мапой документов и количество вхождений этого слова
2й - обработка запроса, где вычленяем уникальные входдения слов и составляем ролевантность документов исходя из совпаения 1го этапа (суммируем к документам входения)
Далее сортируем и выводим первые 5 документов (у которых больше 0 ролевантность)

Временная сложность:
O(n*k), где n - количество документов, k - количество слов в документе:
+
O(m * k * n * nlog(n)), где m - число запросов, k - число уникальных слов в запросе (>=m), n - число документов

Пространственная сложность:
O(m+n+d), m для индекса документов, n для агрегирующей мапы, d для отсортированного слайса документов

*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// SearchLimit лимит на количество ролевантных документов
// Комментарий:
// Да, но это константа подразумевает использование "внутри пакета", те private
// Я конечно же переименовал, но это вроде не best practices =)
const SearchLimit = 5

// searchDoc ролевантный документ
type searchDoc struct {
	idx,
	rel int
}

// Search поисковая структура
type Search struct {
	DocIndex map[string]map[int]int
	RelIndex []*searchDoc
	Result   []int
}

func main() {
	var nd, nq, i int

	input, err := getInputData()
	if err != nil {
		showError(err)
	}

	// получаем количество документов
	input.Scan()
	nd, err = strconv.Atoi(input.Text())
	if err != nil {
		return
	}

	s := NewSearch()
	s.InitDocIndex()

	// собираем индексы по документам
	for i = 1; i <= nd; i++ {
		input.Scan()
		s.CreateDocIndex(strings.Split(input.Text(), " "), i)
	}

	// получаем количество запросов
	input.Scan()
	nq, err = strconv.Atoi(input.Text())
	if err != nil {
		return
	}

	// обрабатываем запросы
	for i = 0; i < nq; i++ {
		input.Scan()

		s.InitRelIndex(nd)
		s.InitResultData(SearchLimit)

		// релевантность документа по слову
		s.CreateQueryIndex(strings.Split(input.Text(), " "))

		// получаем результат
		fmt.Println(strings.Trim(fmt.Sprint(s.GetResultQuery()), "[]"))
	}
}

func testMain(docs []string, queries []string) [][]int {
	s := NewSearch()
	s.InitDocIndex()

	for i := 0; i < len(docs); i++ {
		s.CreateDocIndex(strings.Split(docs[i], " "), i+1)
	}

	var res = make([][]int, 0, len(queries))

	for i := 0; i < len(queries); i++ {

		s.InitRelIndex(len(docs))
		s.InitResultData(SearchLimit)

		s.CreateQueryIndex(strings.Split(queries[i], " "))
		res = append(res, s.GetResultQuery())
	}

	return res
}

// NewSearch инициализация основной структуры
func NewSearch() *Search {
	return &Search{}
}

// InitDocIndex инициализация структуры индекса документов
func (s *Search) InitDocIndex() {
	s.DocIndex = make(map[string]map[int]int)
}

// InitResultData инициализация ответа ролевантных докментов
func (s *Search) InitResultData(size int) {
	s.Result = make([]int, 0, size)
}

// InitRelIndex инициализация индекса ролевантных докментов
func (s *Search) InitRelIndex(size int) {
	s.RelIndex = make([]*searchDoc, 0, size)
}

// SortRel сортировка документов
// если релевантности документов совпадают —– то по возрастанию их порядковых номеров в базе
func (s *Search) SortRel() {
	sort.Slice(s.RelIndex, func(i, j int) bool {
		if s.RelIndex[i].rel != s.RelIndex[j].rel {
			return s.RelIndex[i].rel > s.RelIndex[j].rel
		}
		return s.RelIndex[i].idx < s.RelIndex[j].idx
	})
}

// CreateDocIndex обработка слов и построение индекса по документу
func (s *Search) CreateDocIndex(words []string, docID int) {
	for i := 0; i < len(words); i++ {
		s.AddDocIndex(docID, words[i])
	}
}

// AddDocIndex добавление индекса слова
func (s *Search) AddDocIndex(docID int, word string) {
	if _, ok := s.DocIndex[word]; !ok {
		s.DocIndex[word] = map[int]int{docID: 0}
	}
	s.DocIndex[word][docID]++
}

// CreateQueryIndex обработка слов и построение индекса запроса
func (s *Search) CreateQueryIndex(words []string) {

	mapIdx := make(map[int]int)
	uniqueWords := make(map[string]struct{})

	//s.RelIndex = append(s.RelIndex, &searchDoc{docId: n, rel: r})

	for j := 0; j < len(words); j++ {
		// только уникальные слова
		if _, ok := uniqueWords[words[j]]; !ok {
			// если в документах есть такое слово
			if docCnt, find := s.DocIndex[words[j]]; find {
				for k, v := range docCnt {
					if _, yes := mapIdx[k]; !yes {
						s.RelIndex = append(s.RelIndex, &searchDoc{idx: k, rel: v})
						mapIdx[k] = len(s.RelIndex) - 1
					} else {
						s.RelIndex[mapIdx[k]].rel += v
					}
				}
				uniqueWords[words[j]] = struct{}{}
			}
		}
	}
}

// GetResultQuery результат первых 5ти документов (если ролевантность больше 0)
func (s *Search) GetResultQuery() []int {
	for k, d := range s.RelIndex {
		if k >= SearchLimit || d.rel == 0 {
			break
		}
		s.Result = append(s.Result, d.idx)
	}

	return s.Result
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

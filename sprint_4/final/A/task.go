package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// SearchLimit лимит на количество ролевантных документов
const SearchLimit = 5

// SearchDoc ролевантный документ
type SearchDoc struct {
	id,
	cnt int
}

// Search поисковая структура
type Search struct {
	DocIndex   map[string]map[int]int
	QueryIndex []*SearchDoc
	Result     []int
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

	s := &Search{
		DocIndex: make(map[string]map[int]int),
	}

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

		s.InitRelIndex(nq)
		s.InitResultData(SearchLimit)

		// релевантность документа по слову
		s.CreateQueryIndex(strings.Split(input.Text(), " "))

		// получаем результат
		fmt.Println(strings.Trim(fmt.Sprint(s.GetResultQuery()), "[]"))
	}
}

func testMain(docs []string, queries []string) [][]int {
	s := &Search{
		DocIndex: make(map[string]map[int]int),
	}

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

// InitResultData инициализация ответа ролевантных докментов
func (s *Search) InitResultData(size int) {
	s.Result = make([]int, 0, size)
}

// InitRelIndex инициализация индекса ролевантных докментов
func (s *Search) InitRelIndex(size int) {
	s.QueryIndex = make([]*SearchDoc, 0, size)
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

	for j := 0; j < len(words); j++ {
		// только уникальные слова
		if _, ok := uniqueWords[words[j]]; !ok {
			// если в документах есть такое слово
			if docCnt, find := s.DocIndex[words[j]]; find {
				for k, v := range docCnt {
					if _, yes := mapIdx[k]; !yes {
						s.QueryIndex = append(s.QueryIndex, &SearchDoc{id: k, cnt: v})
						mapIdx[k] = len(s.QueryIndex) - 1
					} else {
						s.QueryIndex[mapIdx[k]].cnt += v
					}
				}
				uniqueWords[words[j]] = struct{}{}
			}
		}
	}

	// сортируем
	s.SortIndex()
}

// GetResultQuery результат первых 5ти документов (если ролевантность больше 0)
func (s *Search) GetResultQuery() []int {
	for i := 0; i < len(s.QueryIndex); i++ {
		if i >= SearchLimit {
			break
		}
		s.Result = append(s.Result, s.QueryIndex[i].id)
	}

	return s.Result
}

// SortIndex сортировка документов (пузырьком)
func (s *Search) SortIndex() {
	if len(s.QueryIndex) == 0 {
		return
	}

	for i := 0; i < len(s.QueryIndex); i++ {
		// больше 5ти не требуется
		if i > SearchLimit {
			break
		}

		// гоним максимальным значением вверх
		for j := len(s.QueryIndex) - 1; j > i; j-- {
			if s.QueryIndex[j-1].cnt < s.QueryIndex[j].cnt ||
				(s.QueryIndex[j-1].cnt == s.QueryIndex[j].cnt && s.QueryIndex[j].id < s.QueryIndex[j-1].id) {
				s.SwapResult(j-1, j)
			}
		}
	}
}

// SwapResult идеоматический свап
func (s *Search) SwapResult(i, j int) {
	s.QueryIndex[i], s.QueryIndex[j] = s.QueryIndex[j], s.QueryIndex[i]
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

package main

import (
	"fmt"
	"reflect"
	"testing"
)

type task struct {
	Input struct {
		Docs,
		Query []string
	}
	Result [][]int
}

func TestTask(t *testing.T) {
	tasks := generateTasks()
	for i := 0; i < len(tasks); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := testMain(tasks[i].Input.Docs, tasks[i].Input.Query)
			if !reflect.DeepEqual(res, tasks[i].Result) {
				t.Errorf("Неверный ответ решения!\nОтвет: %v \nВерно: %v", res, tasks[i].Result)
			}
		})
	}
}

func generateTasks() []*task {
	return []*task{
		{
			Input: struct{ Docs, Query []string }{
				Docs: []string{
					"i love coffee",
					"coffee with milk and sugar",
					"free tea for everyone",
				},
				Query: []string{
					"i like black coffee without milk",
					"everyone loves new year",
					"mary likes black coffee without milk",
				},
			},
			Result: [][]int{
				{1, 2},
				{3},
				{2, 1},
			},
		},
		{
			Input: struct{ Docs, Query []string }{
				Docs: []string{
					"buy flat in moscow",
					"rent flat in moscow",
					"sell flat in moscow",
					"want flat in moscow like crazy",
					"clean flat in moscow on weekends",
					"renovate flat in moscow",
				},
				Query: []string{
					"flat in moscow for crazy weekends",
				},
			},
			Result: [][]int{
				{4, 5, 1, 2, 3},
			},
		},
		{
			Input: struct{ Docs, Query []string }{
				Docs: []string{
					"i like dfs and bfs",
					"i like dfs dfs",
					"i like bfs with bfs and bfs",
				},
				Query: []string{
					"dfs dfs dfs dfs bfs",
				},
			},
			Result: [][]int{
				{3, 1, 2},
			},
		},
	}
}

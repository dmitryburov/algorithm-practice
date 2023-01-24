package main

import (
	"fmt"
	"testing"
)

type task struct {
	N      int
	Arr    []User
	Result []string
}

func TestTask(t *testing.T) {
	tasks := generateTasks()
	for i := 0; i < len(tasks); i++ {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			quickSort(tasks[i].Arr, 0, tasks[i].N-1)

			var strNames = make([]string, tasks[i].N)
			for idx := 0; idx < tasks[i].N; idx++ {
				strNames[idx] = tasks[i].Arr[idx].Login
			}

			if fmt.Sprint(strNames) != fmt.Sprint(tasks[i].Result) {
				t.Errorf("Неверный ответ решения!\nОтвет: %v \nВерно: %v", strNames, tasks[i].Result)
			}
		})
	}
}

func generateTasks() []task {
	return []task{
		{
			N: 5,
			Arr: []User{
				{
					Login:   "alla",
					Success: 4,
					Fail:    100,
				},
				{
					Login:   "gena",
					Success: 6,
					Fail:    1000,
				},
				{
					Login:   "gosha",
					Success: 2,
					Fail:    90,
				},
				{
					Login:   "rita",
					Success: 2,
					Fail:    90,
				},
				{
					Login:   "timofey",
					Success: 4,
					Fail:    80,
				},
			},
			Result: []string{"gena", "timofey", "alla", "gosha", "rita"},
		},
		{
			N: 5,
			Arr: []User{
				{
					Login:   "alla",
					Success: 0,
					Fail:    0,
				},
				{
					Login:   "gena",
					Success: 1,
					Fail:    0,
				},
				{
					Login:   "gosha",
					Success: 2,
					Fail:    0,
				},
				{
					Login:   "rita",
					Success: 3,
					Fail:    0,
				},
				{
					Login:   "timofey",
					Success: 4,
					Fail:    0,
				},
			},
			Result: []string{"timofey", "rita", "gosha", "gena", "alla"},
		},
		{
			N: 5,
			Arr: []User{
				{
					Login:   "alla",
					Success: 0,
					Fail:    0,
				},
				{
					Login:   "allab",
					Success: 0,
					Fail:    0,
				},
				{
					Login:   "allac",
					Success: 0,
					Fail:    0,
				},
				{
					Login:   "allad",
					Success: 0,
					Fail:    0,
				},
				{
					Login:   "allae",
					Success: 0,
					Fail:    0,
				},
			},
			Result: []string{"alla", "allab", "allac", "allad", "allae"},
		},
		{
			N: 2,
			Arr: []User{
				{
					Login:   "za",
					Success: 0,
					Fail:    0,
				},
				{
					Login:   "b",
					Success: 0,
					Fail:    0,
				},
			},
			Result: []string{"b", "za"},
		},
		{
			N: 13,
			Arr: []User{
				{
					Login:   "tufhdbi",
					Success: 76,
					Fail:    58,
				},
				{
					Login:   "rqyoazgbmv",
					Success: 59,
					Fail:    78,
				},
				{
					Login:   "qvgtrlkmyrm",
					Success: 35,
					Fail:    27,
				},
				{
					Login:   "tgcytmfpj",
					Success: 70,
					Fail:    27,
				},
				{
					Login:   "xvf",
					Success: 84,
					Fail:    19,
				},
				{
					Login:   "jzpnpgpcqbsmczrgvsu",
					Success: 30,
					Fail:    3,
				},
				{
					Login:   "evjphqnevjqakze",
					Success: 92,
					Fail:    15,
				},
				{
					Login:   "wwzwv",
					Success: 87,
					Fail:    8,
				},
				{
					Login:   "tfpiqpwmkkduhcupp",
					Success: 1,
					Fail:    82,
				},
				{
					Login:   "tzamkyqadmybky",
					Success: 5,
					Fail:    81,
				},
				{
					Login:   "amotrxgba",
					Success: 0,
					Fail:    6,
				},
				{
					Login:   "easfsifbzkfezn",
					Success: 100,
					Fail:    28,
				},
				{
					Login:   "kivdiy",
					Success: 70,
					Fail:    47,
				},
			},
			Result: []string{"easfsifbzkfezn", "evjphqnevjqakze", "wwzwv", "xvf", "tufhdbi", "tgcytmfpj", "kivdiy", "rqyoazgbmv", "qvgtrlkmyrm", "jzpnpgpcqbsmczrgvsu", "tzamkyqadmybky", "tfpiqpwmkkduhcupp", "amotrxgba"},
		},
	}
}

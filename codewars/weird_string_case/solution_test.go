package main

import (
	"fmt"
	"reflect"
	"testing"
)

type TestItem struct {
	input  string
	output string
}

func TestTask(t *testing.T) {
	for i, v := range generateTasks() {
		t.Run(fmt.Sprintf("Test %d", i+1), func(t *testing.T) {
			res := toWeirdCase(v.input)
			if !reflect.DeepEqual(res, v.output) {
				t.Errorf("Wrong test.\nOutput: \n%v \nExpected: \n%v", res, v.output)
			}
		})
	}
}

func generateTasks() []TestItem {
	return []TestItem{
		{
			input:  "abc def",
			output: "AbC DeF",
		},
		{
			input:  "ABC",
			output: "AbC",
		},
		{
			input:  "This is a test Looks like you passed",
			output: "ThIs Is A TeSt LoOkS LiKe YoU PaSsEd",
		},
		{
			input:  "Yo",
			output: "Yo",
		},
	}
}

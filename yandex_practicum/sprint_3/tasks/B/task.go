package main

import (
	"fmt"
	"strings"
)

var numberRules = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func main() {
	var input string
	_, _ = fmt.Scan(&input)

	fmt.Println(strings.Trim(fmt.Sprint(solution(input)), "[]"))
}

func solution(digits string) []string {
	if digits == "" {
		return nil
	}

	var result []string
	var words = numberRules[digits[len(digits)-1:]]
	var combinations = solution(digits[:len(digits)-1])

	if len(combinations) > 0 {
		for _, combination := range combinations {
			for _, c := range strings.Split(words, "") {
				result = append(result, combination+c)
			}
		}
	} else {
		for _, c := range strings.Split(words, "") {
			result = append(result, c)
		}
	}

	return result
}

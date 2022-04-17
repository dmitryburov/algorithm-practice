package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	str1, str2, err := getInputData()
	if err != nil {
		showError(err)
	}

	fmt.Println(solution(str1, str2))
}

func solution(s, t string) string {
	lenS, lenT := len(s), len(t)

	if lenS == 0 || s == " " {
		return "True"
	}

	if lenS > lenT {
		return "False"
	}

	for i, j := 0, 0; j < lenT; j++ {
		if s[i] == t[j] {
			i++
			if i == lenS {
				return "True"
			}
		}
	}

	return "False"
}

func getInputData() (s, t string, err error) {
	var input *os.File
	var btStr []byte

	input, err = getInputFromFile()
	if err != nil {
		showError(err)
	}
	// close file
	defer func(input *os.File) {
		_ = input.Close()
	}(input)

	reader := bufio.NewReader(input)

	btStr, _, err = reader.ReadLine()
	if err != nil {
		return
	}
	s = string(btStr)

	btStr, _, err = reader.ReadLine()
	if err != nil {
		return
	}
	t = string(btStr)

	// clear bufio
	defer reader.Reset(reader)
	return
}

func getInputFromFile() (*os.File, error) {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil, err
	}

	return file, nil
}

func showError(err interface{}) {
	panic(err)
}

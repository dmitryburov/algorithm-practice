package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// F. Палиндром

func main() {
	var inp = GetInputFromFile()
	defer inp.Close()

	reader := bufio.NewReader(inp)
	str, _ := reader.ReadString('\n')

	str = strings.ToLower(str)
	re := regexp.MustCompile("[^a-z]")
	str = re.ReplaceAllString(str, "")

	var result = "True"
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-(i+1)] {
			result = "False"
			break
		}
	}

	fmt.Println(result)
}

func GetInputFromFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	return file
}

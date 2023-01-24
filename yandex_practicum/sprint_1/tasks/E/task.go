package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inp := GetInputFromFile()
	defer inp.Close()

	var find string

	reader := bufio.NewReader(inp)
	_, _ = reader.ReadString('\n')
	line, _ := reader.ReadString('\n')
	strItems := strings.Split(strings.Trim(line, ""), " ")

	for i := 0; i < len(strItems); i++ {
		trimStr := strings.TrimSpace(strItems[i])
		if len(trimStr) > len(find) {
			find = trimStr
		}
	}

	fmt.Println(find)
	fmt.Println(len(find))
}

func GetInputFromFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	return file
}

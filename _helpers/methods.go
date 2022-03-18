package _helpers

import (
	"bufio"
	"io"
	"os"
)

func GetInputFromFile() *os.File {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}

	return file
}

func GetInputFromConsole() io.Reader {
	return bufio.NewReader(os.Stdin)
}

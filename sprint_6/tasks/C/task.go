package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Color uint8

const (
	_ Color = iota
	WHITE
	GRAY
	BLACK
)

type Node struct {
	value int
	to    []*Node
}

type NodeList []*Node

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {

}

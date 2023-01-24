package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type GoldenHeap struct {
	price int
	num   int
}

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)

	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {

	var totalCost int

	m, k, data := getData(bufio.NewScanner(r))

	data = merge(data)
	for i := 0; i < k && m > 0; i++ {
		for g := 0; g < data[i].num && m > 0; g++ {
			totalCost += data[i].price
			m--
		}
	}

	s.WriteString(fmt.Sprint(totalCost))
}

func getData(scanner *bufio.Scanner) (m, k int, items []GoldenHeap) {
	var err error

	scanner.Scan()
	m, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	k, err = strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	items = make([]GoldenHeap, k)

	var a, b int
	for i := 0; i < k; i++ {
		scanner.Scan()
		arrStr := strings.Split(scanner.Text(), " ")

		a, err = strconv.Atoi(arrStr[0])
		if err != nil {
			log.Fatal(err)
		}
		b, err = strconv.Atoi(arrStr[1])
		if err != nil {
			log.Fatal(err)
		}

		items[i] = GoldenHeap{
			price: a,
			num:   b,
		}
	}

	return
}

func merge(mine []GoldenHeap) []GoldenHeap {
	mineLength := len(mine)
	if mineLength < 2 {
		return mine
	}

	mid := mineLength >> 1

	a := merge(mine[:mid])
	aLength := len(a)

	b := merge(mine[mid:])
	bLength := len(b)

	var aIndex, bIndex int
	result := make([]GoldenHeap, 0, mineLength)

	for aIndex < aLength || bIndex < bLength {
		if aIndex == aLength {
			result = append(result, b[bIndex])
			bIndex++
			continue
		}
		if bIndex == bLength {
			result = append(result, a[aIndex])
			aIndex++
			continue
		}
		if a[aIndex].price >= b[bIndex].price {
			result = append(result, a[aIndex])
			aIndex++
		} else {
			result = append(result, b[bIndex])
			bIndex++
		}
	}

	return result

}

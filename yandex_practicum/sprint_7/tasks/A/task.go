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

type Price struct {
	day         int
	price       int
	currentPick bool
}

func main() {
	s := strings.Builder{}
	Solution(os.Stdin, &s)
	fmt.Println(s.String())
}

func Solution(r io.Reader, s *strings.Builder) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()

	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	priceArr := strings.Fields(scanner.Text())

	priceData := make([]Price, 0, n)
	for i, priceStr := range priceArr {
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			log.Fatal(err)
		}
		priceData = append(priceData, Price{
			day:   i,
			price: price,
		})
	}

	var lowPick Price
	value := 0
	for i := 0; i < n; i++ {
		if !lowPick.currentPick {
			if i+1 < n &&
				priceData[i+1].price < priceData[i].price {
				continue
			}
			lowPick = priceData[i]
			lowPick.currentPick = true
			continue
		}
		if i+1 < n && priceData[i].price < priceData[i+1].price {
			continue
		}
		value += priceData[i].price - lowPick.price
		lowPick.currentPick = false
	}

	s.WriteString(fmt.Sprint(value))
}

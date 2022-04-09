package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// E. Самое длинное слово
// Чтобы подготовиться к семинару, Гоше надо прочитать статью по эффективному менеджменту.
// Так как Гоша хочет спланировать день заранее, ему необходимо оценить сложность статьи.
// Он придумал такой метод оценки: берётся случайное предложение из текста и в нём ищется самое длинное слово.
// Его длина и будет условной сложностью статьи. Помогите Гоше справиться с этой задачей.

// Формат ввода:
// В первой строке дана длина текста L (1 ≤ L ≤ 10(5)).
// В следующей строке записан текст, состоящий из строчных латинских букв и пробелов. Слово —– последовательность букв, не разделённых пробелами. Пробелы могут стоять в самом начале строки и в самом её конце.
// Текст заканчивается переносом строки, этот символ не включается в число остальных L символов.
//
// Формат вывода:
// В первой строке выведите самое длинное слово. Во второй строке выведите его длину.
// Если подходящих слов несколько, выведите то, которое встречается раньше.
//
// Примеры:
// inp: 					>>> 	out:
//    19								segment
//    i love segment tree				7
//
// inp: 					>>> 	out:
//    21								jumps
//	  frog jumps from river				5
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

package main

import (
	"strings"
	"unicode"
)

func toWeirdCase(str string) string {
	var result strings.Builder

	i := 1
	for _, ch := range str {
		if unicode.IsSpace(ch) || !unicode.IsLetter(ch) {
			i = 1
			result.WriteRune(ch)
		} else {
			if i%2 == 0 {
				result.WriteByte(byte(unicode.ToLower(ch)))
			} else {
				result.WriteByte(byte(unicode.ToUpper(ch)))
			}
			i++
		}
	}

	return result.String()
}

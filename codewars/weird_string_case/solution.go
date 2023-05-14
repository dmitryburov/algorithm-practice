package main

import "unicode"

func toWeirdCase(str string) string {
	i := 1
	var result = make([]rune, len(str))

	for idx, ch := range str {
		switch {
		case unicode.IsSpace(ch):
			i = 1
			result[idx] = ch
		case unicode.IsLetter(ch):
			if i%2 == 0 {
				result[idx] = unicode.ToLower(ch)
			} else {
				result[idx] = unicode.ToUpper(ch)
			}
			i++
		}
	}

	return string(result)
}

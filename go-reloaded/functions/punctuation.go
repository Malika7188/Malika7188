package goreloaded

import (
	"strings"
	"unicode"
)

func Punc(st []string) []string {
	s := strings.Join(st, " ")
	nRune := []rune(s)
	var result []rune
	for i, val := range nRune {
		if val != ' ' && i < len(nRune) && !(unicode.IsLetter(val)) && !(unicode.IsDigit(val)) && val != '\'' {
			if nRune[i-1] == ' ' {
				nRune[i], nRune[i-1] = nRune[i-1], nRune[i]
			}
		}
	}
	for i, char := range nRune {
		if i < len(nRune)-1 && char == ' ' && nRune[i-1] == ' ' {
			continue
		}
		result = append(result, nRune[i])
	}
	return strings.Split(string(result), " ")
}

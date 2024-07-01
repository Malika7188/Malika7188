package functions

import "strings"

func Backspace(s string) string {
	index := strings.Index(s, "\\b")

	if index == -1 {
		return s
	}
	if index == 0 {
		return Backspace(s[1:])
	}
	return Backspace(s[:index-1] + s[index+2:])
}
